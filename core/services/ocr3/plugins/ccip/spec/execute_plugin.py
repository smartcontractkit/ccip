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

class ExecuteObservation:
    reports map[ChainSelector][]ExecutePluginCommitData
    messages map[ChainSelector]map[ChainSelector]CCIPMsg

class ExecuteOutcome:
    report_data []ExecutePluginCommitDataWithMessages
    pending_reports []ExecutePluginCommitDataWithMessages

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

    def observation_quorum(self):
        return "F+1"

    def outcome(self, previous_outcome, observations):
        # merge observations, removing any which do not reach f_chain threshold.
        commit_reports = merge_commit_observations(observations, self.cfg["f_chain"])
        messages = merge_message_observations(observations, self.cfg["f_chain"])

        # flatten report map and sort by timestamp
        flattened_reports = flattenReports(commit_reports)

        # flatten messages map and sort by report timestamp/sequence number
        flattened_messages = []
        for report in flattened_reports:
            for i in (report.seq_num_range.start, report.seq_num_range.end+1):
                if i in messages[report.chain]:
                    flattened_messages.append(messages[report.chain][i])


        # merge new observations with pending data from previous outcome
        pending_reports = mergeReports(flattened_messages, previous_outcome.pending_reports, messages)

        # select reports from pending data to include in the final report
        report_data, pending_reports = selectReportData(pending_reports)

        return (report_data, pending_reports)

    def reports(self, outcome):
        report = report_from_outcome(outcome)
        encoded = report.chain_encode() # abi_encode for evm chains
        return [encoded]

    def should_accept(self, report):

    def should_transmit(self, report):

    def keep_cfg_in_sync(self):
        # Polling the configuration on the on-chain contract.
        # When the config is updated on-chain, updates the plugin's local copy to the most recent version.
        pass
