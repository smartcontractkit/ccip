
# This is a high level description of the commit plugin inner workings.

class CommitPlugin:
    def __init__(self):
        self.cfg = {
            "dest_chain": "chainD",
            "f_chain": {
                "chainA": 2,
                "chainB": 3,
                "chainD": 3
            },
            "observer_info": {
                "nodeA": {
                    "supported_chains": {"chainA", "chainB", "chainD"},
                    "token_prices_observer": True,
                }
            },
            "priced_tokens": {"tokenA", "tokenB"}
        }

        self.offRamp = {}
        self.onRamps = []
        self.known_source_chains = set() # Initially populated based on all the chains of the config.

    def query(self):
        pass

    def observation(self, previous_outcome):
        # Observe sequence numbers.
        observed_seq_nums = {}
        if previous_outcome:
            observed_seq_nums = previous_outcome["observed_seq_nums"]
        if self.can_read_dest():
            for (chain, seq_num) in self.offRamp.last_sequence_numbers():
                if chain not in observed_seq_nums or seq_num > observed_seq_nums[chain]:
                    observed_seq_nums[chain] = seq_num

        # Observe new msgs. (id, seq_num)
        new_msgs = {} # chain -> [(id, seq_num)]
        for (chain, seq_num) in observed_seq_nums.items():
            if not self.can_read(chain, seq_num):
                continue
            new_msgs[chain] log_poller.get_new_msgs(chain, from=seq_num+1, limit=256)

        # Observe token prices.
        token_prices = self.get_token_prices()

        # Observe gas prices.
        gas_prices = self.get_gas_prices()

        return {
            "seq_nums": observed_seq_nums,
            "msgs": new_msgs,
            "token_prices": token_prices,
            "gas_prices": gas_prices,
            "f_chain": self.cfg["f_chain"]
        }

    def validate_observation(self, observation):
        for (chain, msg) in observation["msgs"].items():
            # Message sequence number must be greater than the reported max sequence number for that chain.

            # Message id and sequence number must be unique inside this observation.

            # Observer must be able to read the chain.

        for (token, price) in observation["token_prices"]:
            # Ensure no duplicate price for the same token.

        for (chain, gas_price) in observation["gas_prices"]:
            # Ensure no duplicate gas price for the same chain.

    def observation_quorum(self):
        return "2F+1"

    def outcome(self, observations):
        # For each chain:

        # Keep the most voted fChain.

        # Require 2*fChain+1 sequence number observations.
        # Order the sequence numbers and take sequence_number[fChain].

        # Filter out messages with sequence number less than consensus sequence number.
        # Assume for a sequence number that the msg id is the one with the most votes (require 2*fChain+1 votes).
        # After that nodes agree about a set of msgs ["1:0x1", "2:0x2", "4:0x4"]
        # Starting from the first one build the root until finding a gap, in this example [1,2].
        # Build merkle tree using leaves ["0x1", "0x2"].

        # For each token:
        # Take median price.

        # For each gas price:
        # Take median gas price.

    def reports(self, outcome):
        # Simply builds and encodes a report aimed for destination chain.

    def should_accept(self, report):
        if report is empty or invalid:
            return False

    def should_transmit(self, report):
        if not self.is_writer():
            return False

        if report is empty or invalid:
            return False

        if report contains at least one sequence number that is already committed:
            return False # Since it will revert on-chain.

        return True
