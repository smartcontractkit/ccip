#
# High-level Python specification for the CCIP OCR3 Execute Plugin.
#
# This specification aims to provide a clear and comprehensive understanding
# of the plugin's functionality. It is highly recommended for engineers working on CCIP
# to familiarize themselves with this specification prior to reading the
# corresponding Go implementation.
#
# NOTE: Even though the specification is written in a high-level programming language, it's purpose
# is not to be executed. It is meant to be just a reference for the Go implementation.
#

from typing import Dict, List, Set

type ChainSelector = int
type SequenceNumber = int

class ExecutePluginCommitData:
    pass

class CCIPMsg:
    pass

class ExecuteObservation:
    reports:   Dict[ChainSelector, List[ExecutePluginCommitData]]
    messages:  Dict[ChainSelector, Dict[SequenceNumber, CCIPMsg]]

class ExecuteOutcome:
    report_data:     List[ExecutePluginCommitDataWithMessages]
    pending_reports: List[ExecutePluginCommitDataWithMessages]

class ExecutePlugin:
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
         }
         self.keep_cfg_in_sync()

    def query(self):
        pass

    def observation(self, previous_outcome):
        # Observe commit reports
        if self.can_read_dest():
            pending_reports = pending_reports + self.offRamp.get_pending_commit_reports(self.latest_report)
            self.latest_report = pendingReport[-1].timestamp

        # Observe messages from previous outcome
        for (selector, reports) in previous_outcome.pending_reports:
            for report in reports:
                messages = self.onRamp.get_messages(report.seq_num_range)
                for (seq_num, msg) in messages:
                    messages[selector][seq_num] = msg

        return (reports, messages)


    def validate_observation(self, attributed_observation):

        # Validate observations are allowed by configuration
        observation = attributed_observation.observation
        observer = attributed_observation.observer

        if "reports" in observation:
            assert observer.can_read_dest()

        observer_supported_chains = self.cfg["observer_info"][observer]["supported_chains"]
        for (chain, msgs) in observation["new_msgs"].items():
            assert(chain in observer_supported_chains)

        # validate reports do not contain duplicates or overlapping ranges
        for (chain, reports) in observation["reports"].items():
            roots = Set[string]
            ranges = List[SeqNumRange]

          for report in reports:
            assert(not report.root.string() in roots.values)
            roots.add(report.root.string())

            for r in ranges:
                assert(not seq_range_overlaps(report.seqNumRange, r))

            for seqNum in report.ExecutedMessages:
                assert(seqNum >= report.seqNumRange.start && seqNum <= report.seqNumRange.end)

    def observation_quorum(self):
        return "F+1"

    def select_report_data(reports_with_messages):
        """select oldest (front of list) reports and build as many proofs as we
        can fit for the execute report"""
        proofs = []
        proofs_bytes = 0
        index = 0
        for report in reports_with_messages:
            # build a report to execute as many messages as will fit in the
            # remaining transmit size.
            proof, fully_execute = report.build_proof(MAX_REPORT_SIZE - report_data_bytes)
            if not proof.empty():
                proofs.append(proof)

            if fully_execute:
                index++
            else:
                return proofs, reports_with_messages[index:]

    def outcome(self, previous_outcome, observations):
        # merge observations, removing any which do not reach f_chain threshold.
        commit_reports = merge_commit_observations(observations, self.cfg["f_chain"])
        messages = merge_message_observations(observations, self.cfg["f_chain"])

        # flatten report map and sort by timestamp/sequence number
        flattened_reports = flatten_reports(commit_reports)

        # add messages to report object
        for report in flattened_reports:
            for i in (report.seq_num_range.start, report.seq_num_range.end+1):
                if i in messages[report.chain]:
                    report.messages.append(messages[report.chain][i])


        # take the new flattened reports and merge with the previous pending reports
        pending_reports = merge_reports(flattened_messages, previous_outcome.pending_reports)

        # select reports from pending data to include in the final report
        report_data, pending_reports = select_report_data(pending_reports)

        return (report_data, pending_reports)

    def reports(self, outcome):
        report = report_from_outcome(outcome)
        encoded = report.chain_encode() # abi_encode for evm chains
        return [encoded]

    def should_accept(self, report):
        pass

    def should_transmit(self, report):
        pass

    def keep_cfg_in_sync(self):
        # Polling the configuration on the on-chain contract.
        # When the config is updated on-chain, updates the plugin's local copy to the most recent version.
        pass
