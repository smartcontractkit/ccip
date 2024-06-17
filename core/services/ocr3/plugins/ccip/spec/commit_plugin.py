#
# High-level Python specification for the CCIP OCR3 Commit Plugin.
#
# This specification aims to provide a clear and comprehensive understanding
# of the plugin's functionality. It is highly recommended for engineers working on CCIP
# to familiarize themselves with this specification prior to reading the
# corresponding Go implementation.
#
# NOTE: Even though the specification is written in a high-level programming language, it's purpose
# is not to be executed. It is meant to be just a reference for the Go implementation.
#
class CommitPlugin:
    def __init__(self):
         self.cfg = {
             "dest_chain": "chainD",
             "f_chain": {"chainA": 2, "chainB": 3, "chainD": 3},
             "observer_info": {
                 "nodeA": {
                     "supported_chains": {"chainA", "chainB", "chainD"},
                     "token_prices_observer": True,
                 }
             },
             "priced_tokens": {"tokenA", "tokenB"},
         }
         self.keep_cfg_in_sync()

    def query(self):
        pass

    def observation(self, previous_outcome):
        # Observe last msg sequence numbers for each source chain: {sourceChain: sequenceNumber}
        observed_seq_nums = previous_outcome.get("observed_seq_nums", default={})
        if self.can_read_dest():
            on_chain_seq_nums = self.offRamp.get_sequence_numbers()
            for (chain, seq_num) in on_chain_seq_nums.items():
                if chain not in observed_seq_nums or seq_num > observed_seq_nums[chain]:
                    observed_seq_nums[chain] = seq_num

        # Observe new msgs: {sourceChain: [(seq_num, hash)]}
        new_msgs = {}
        for (chain, seq_num) in observed_seq_nums.items():
            if self.can_read(chain):
                new_msgs[chain] = self.onRamp(chain).get_msgs(chain, start=seq_num+1, limit=256)

        # Observe token prices. {token: price}
        token_prices = self.get_token_prices()

        # Observe gas prices. {chain: gasPrice}
        gas_prices = self.get_gas_prices()

        # Observe fChain for each chain. {chain: f_chain}
        f_chain = self.cfg["f_chain"]

        if not self.can_read_dest():
            # If node is not able to read updated sequence numbers from the destination,
            # it should not observe the outdated ones that are coming from the previous outcome.
            observed_seq_nums = {}

        return (observed_seq_nums, new_msgs, token_prices, gas_prices, f_chain)


    def validate_observation(self, attributed_observation):
        observation = attributed_observation.observation
        observer = attributed_observation.observer

        if "seq_nums" in observation:
            assert observer.can_read_dest()

        observer_supported_chains = self.cfg["observer_info"][observer]["supported_chains"]
        for (chain, msgs) in observation["new_msgs"].items():
            assert(chain in observer_supported_chains)

            if "seq_nums" in observation:
                for msg in msgs:
                    assert(msg.seq_num > observation["observed_seq_nums"][msg.source_chain])

            assert(len(msgs) == len(set([msg.seq_num for msg in msgs])))
            assert(len(msgs) == len(set([msg.hash for msg in msgs])))

    def observation_quorum(self):
        return "2F+1"

    def outcome(self, observations):
        f_chain = consensus_f_chain(observations)
        seq_nums = consensus_seq_nums(observations, f_chain)

        # all_msgs contains all messages from all observations, grouped by source chain
        all_msgs = [observation["new_msgs"] for observation in observations].group_by_source_chain()

        trees = {} # { chain: (root, min_seq_num, max_seq_num) }
        for (chain, msgs) in all_msgs:
            # keep only msgs with seq nums greater than the consensus max commited seq nums
            msgs = [msg for msg in msgs if msg.seq_num > seq_nums[chain]]

            msgs_by_seq_num = msgs.group_by_seq_num() # { 423: [0x1, 0x1, 0x2] }
                                                      # 2 nodes say that msg hash is 0x1 and 1 node says it's 0x2

            msg_hashes = { seq_num: elem_most_occurrences(hashes) for (seq_num, hashes) in msgs_by_seq_num.items() }
            for (seq_num, hash) in msg_hashes.items(): # require at least 2f+1 observations of the voted hash
                assert(msgs_by_seq_num[seq_num].count(hash) >= 2*f_chain[chain]+1)

            msgs_for_tree = [] # [ (seq_num, hash) ]
            for (seq_num, hash) in msg_hashes.ordered_by_seq_num():
                if len(msgs_for_tree) > 0 and msgs_for_tree[-1].seq_num+1 != seq_num:
                    break # gap in sequence numbers, stop here
                msgs_for_tree.append((seq_num, hash))

            trees[chain] = build_merkle_tree(msgs_for_tree, leaves="hashes")

        token_prices = { tk: median(prices) for (tk, prices) in observations.group_token_prices_by_token() }
        gas_prices = { chain: median(prices) for (chain, prices) in observations.group_gas_prices_by_chain() }

        return (seq_nums, trees, token_prices, gas_prices)

    def reports(self, outcome):
        report = report_from_outcome(outcome)
        encoded = report.chain_encode() # abi_encode for evm chains
        return [encoded]

    def should_accept(self, report):
        if report is empty or invalid:
            return False

    def should_transmit(self, report):
        if not self.is_writer():
            return False

        if report is empty or invalid:
            return False

        on_chain_seq_nums = self.offRamp.get_sequence_numbers()
        for (chain, tree) in report.trees():
            if not (on_chain_seq_nums[chain]+1 == tree.min_seq_num):
                return False

        return True

    def keep_cfg_in_sync(self):
        # Polling the configuration on the on-chain contract.
        # When the config is updated on-chain, updates the plugin's local copy to the most recent version.
        pass

def consensus_f_chain(observations):
    f_chain_votes = observations["f_chain"].group_by_chain() # { chainA: [1, 1, 16, 16, 16, 16] }
    return { ch: elem_most_occurrences(fs) for (ch, fs) in f_chain_votes.items() } # { chainA: 16 }

def consensus_seq_nums(observations, f_chain):
    observed_seq_nums = observations["observed_seq_nums"].group_by_chain(sort="asc") # { chainA: [4, 5, 5, 5, 5, 6, 6] }
    seq_nums_consensus = {}

    for chain, seq_nums in observed_seq_nums.items():
        if len(observed_seq_nums) >= 2*f_chain[chain]+1:
            seq_nums_consensus[chain] = observed_seq_nums[f_chain[chain]] # with f=4 { chainA: 5 }

    return seq_nums_consensus
