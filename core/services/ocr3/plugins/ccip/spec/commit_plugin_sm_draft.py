#
# High-level Python specification for the CCIPv2 OCR3 Commit Plugin, with a focus on RMN OffChain Blessing
#
# The plugin is implemented as a state machine, and moves from state to state each round. There are 3 states:
# 1. Determining intervals to be used in the next round
# 2. Building a report from the intervals determined in the previous round
# 3. Checking if the most recently generated report has been committed to the dest chain
#    - if the report has been committed, move to state 1
#    - else, repeat step 3
#
# This approach leads to a clear separation of concerns and addresses the complications that can arise if a report
# is not successfully transmitted (as we explicitly only continue once we know the previous report has been committed).


from typing import List, Dict, Optional
from dataclasses import dataclass
from collections import defaultdict

RmnNodeId = str
ChainSelector = int

MAX_INTERVAL_LENGTH = 256


@dataclass
class Interval:
    min: int
    max: int

    def is_empty(self) -> bool:
        return self.min == self.max


@dataclass
class RmnSig:
    rmn_node_id: RmnNodeId
    sig: bytes


@dataclass
class SignedInterval:
    interval: Interval
    root: bytes
    sigs: List[RmnSig]


@dataclass
class CcipMessage:
    seq_num: int


@dataclass
class CommitQuery:
    rmn_max_seq_nums: Dict[ChainSelector, int]
    signed_intervals: Dict[ChainSelector, SignedInterval]


@dataclass
class MaxSourceChainSeqNumsObservation:
    max_source_chain_seq_nums: Dict[ChainSelector, int]


@dataclass
class MerkleRootsObservation:
    merkle_roots: Dict[ChainSelector, bytes]


@dataclass
class MaxCommittedSeqNumsObservation:
    max_committed_seq_nums: Dict[ChainSelector, int]


CommitObservation = MaxSourceChainSeqNumsObservation | MerkleRootsObservation | MaxCommittedSeqNumsObservation


# TODO: doc
@dataclass
class NextIntervalsChosen:
    intervals: Dict[ChainSelector, Interval]


# TODO: doc
@dataclass
class ReportGenerated:
    signed_intervals: Dict[ChainSelector, SignedInterval]


# TODO: doc
@dataclass
class ReportInFlight:
    signed_intervals: Dict[ChainSelector, SignedInterval]
    attempts: int


# TODO: doc
@dataclass
class ReportCommitted:
    max_committed_seq_nums: Dict[ChainSelector, int]


CommitOutcome = NextIntervalsChosen | ReportGenerated | ReportInFlight | ReportCommitted


@dataclass
class CommitReport:
    signed_intervals: Dict[ChainSelector, SignedInterval]


@dataclass
class RmnNode:
    node_id: RmnNodeId
    ip_address: bytes
    pub_key: bytes
    supported_chains: List[ChainSelector]


@dataclass
class RmnClientConfig:
    rmn_nodes: List[RmnNode]


@dataclass
class RmnClient:
    rmn_client_config: RmnClientConfig

    # TODO: doc
    def request_max_seq_nums(
            self,
            chains: List[ChainSelector]
    ) -> Dict[ChainSelector, int]:
        pass

    # TODO: doc
    def request_signed_intervals(
            self,
            intervals: Dict[ChainSelector, Interval]
    ) -> Dict[ChainSelector, SignedInterval]:
        pass


class ChainReader:
    def __init__(self):
        pass


class OffRamp:
    def __init__(self):
        pass

    # TODO: doc
    def get_max_seq_nums_on_dest_chain(self) -> Dict[ChainSelector, int]:
        pass


@dataclass
class CommitPlugin:
    rmn_client: RmnClient
    all_source_chains: List[ChainSelector]
    dest_chain: ChainSelector
    chain_readers: Dict[ChainSelector, ChainReader]
    off_ramp: OffRamp
    f: int
    max_check_report_persisted_attempts: int

    # TODO: doc
    def can_read_from_dest_chain(self) -> bool:
        return self.dest_chain in self.chain_readers

    # TODO: doc
    def get_ccip_messages_from_source_chains(
            self,
            intervals: Dict[ChainSelector, Interval]
    ) -> Dict[ChainSelector, List[CcipMessage]]:
        pass

    # TODO: doc
    def query(self, previous_outcome: CommitOutcome) -> CommitQuery:
        match previous_outcome:
            # If the previous round choose an interval, this round we should build a report from it. We request signed
            # intervals from RMN which need to be included in the report.
            case NextIntervalsChosen(intervals):
                signed_intervals = self.rmn_client.request_signed_intervals(intervals)
                return CommitQuery({}, signed_intervals)

            # If the previous round generated a report or did not detect the report on the dest chain, we need to query
            # the dest chain again to detect if the report has since been committed. In this case we don't need to make
            # a request to RMN.
            case ReportGenerated(_) | ReportInFlight(_, _):
                return CommitQuery({}, {})

            # If in the previous round we detected the report was committed on the dest chain, we need to determine the
            # intervals to use to build a report in the next round. We need to query RMN for the max sequence numbers
            # it has for each source chain, so we can set appropriate upper ranges for our intervals.
            case ReportCommitted(_):
                rmn_max_seq_nums = self.rmn_client.request_max_seq_nums(self.all_source_chains)
                return CommitQuery(rmn_max_seq_nums, {})

    # Given a mapping from source chains to intervals, for each chain read the messages in its interval and compute
    # the merkle root. Return a mapping from chain to merkle root.
    def get_merkle_roots(self, intervals: Dict[ChainSelector, Interval]) -> Dict[ChainSelector, bytes]:
        pass

    # Read from the dest chain (if possible) and return a mapping from source chain to the maximum sequence number that
    # has been committed to the OffRamp on the dest chain.
    def get_max_committed_seq_nums(self) -> Dict[ChainSelector, int]:
        pass

    # For each source chain, return the maximum sequence number (for the dest chain) that is on the source chain's
    # OnRamp
    def get_source_chain_max_seq_nums(self) -> Dict[ChainSelector, int]:
        pass

    # Given intervals that have been reported, and the maximum committed sequence numbers that are actually on chain,
    # return True if reported_intervals have been committed, False otherwise
    def reported_intervals_are_persisted(
            self,
            reported_intervals: Dict[ChainSelector, Interval],
            max_committed_seq_nums: Dict[ChainSelector, int]
    ) -> bool:
        pass

    # The OCR3 implementation of Observation
    def observation(self, previous_outcome: CommitOutcome) -> CommitObservation:
        match previous_outcome:
            # If the previous round choose an interval, this round we should observe the merkle roots of all intervals
            # on all associated source chains
            case NextIntervalsChosen(intervals):
                return MerkleRootsObservation(self.get_merkle_roots(intervals))

            # If the previous round generated a report or did not detect the report on the dest chain, we need to
            # observe the maximum committed sequence numbers on the dest chain, which we will then use to determine
            # if the most recent report has been committed.
            case ReportGenerated(_) | ReportInFlight(_, _):
                return MaxCommittedSeqNumsObservation(self.get_max_committed_seq_nums())

            # If in the previous round we detected the report was committed on the dest chain, we need to determine the
            # intervals to use to build the next report in the next round. Observe the maximum sequence number of
            # messages on the source chains, which we'll use to determine the upper limit of intervals.
            case ReportCommitted(_):
                MaxSourceChainSeqNumsObservation(self.get_source_chain_max_seq_nums())

    # Given a list of MerkleRootObservations, return a flattened consensus on the merkle root for each source chain
    def flatten_merkle_root_observations(self, observations: List[CommitObservation]) -> Dict[ChainSelector, bytes]:
        pass

    # Verify the RMN signatures on the given signed_intervals
    def verify_signed_intervals(
            self,
            signed_intervals: Dict[ChainSelector, SignedInterval]
    ) -> Dict[ChainSelector, SignedInterval]:
        pass

    # Given the signed intervals from RMN (collected in Query) and the merkle roots observed in Observe, return the
    # set of signed intervals that should be included in the report. This essentially combines and reconciles
    # rmn_signed_intervals and observed_merkle_roots. For example, if rmn_signed_intervals and observed_merkle_roots
    # have different merkle roots for the same chain, this chain is not included in the output. Additionally, if there
    # are chains that don't require RMN support, these chains will be in observed_merkle_roots but not
    # rmn_signed_intervals, and will be included in the output (with an empty set of RMN signatures).
    def get_signed_intervals_to_report(
            self,
            intervals: Dict[ChainSelector, Interval],
            rmn_signed_intervals: Dict[ChainSelector, SignedInterval],
            observed_merkle_roots: Dict[ChainSelector, bytes]
    ) -> Dict[ChainSelector, SignedInterval]:
        pass

    # Given a list of MaxCommittedSeqNumsObservation, return a flattened consensus on the max committed sequence number
    # for each source chain
    def flatten_max_committed_seq_nums_observations(
            self,
            observations: List[CommitObservation]
    ) -> Dict[ChainSelector, int]:
        pass

    # Return True if a report with the given signed_intervals has been committed to dest chain (by checking against
    # max_committed_seq_nums), return False otherwise
    def report_has_been_committed(
            self,
            signed_intervals: Dict[ChainSelector, SignedInterval],
            max_committed_seq_nums: Dict[ChainSelector, int]
    ) -> bool:
        pass

    # Given a list of MaxSourceChainSeqNumsObservations, return a flattened consensus on the max sequence number
    # for each source chain
    def flatten_max_source_chain_seq_nums_observations(
            self,
            observations: List[CommitObservation]
    ) -> Dict[ChainSelector, int]:
        pass

    # Given the RMN max sequence numbers for each chain, the max committed sequence numbers on the dest chain, and the
    # max sequence numbers on the chains' OnRamps, return the intervals for the next round.
    def get_next_intervals(
            self,
            max_source_chain_seq_nums: Dict[ChainSelector, int],
            rmn_max_seq_nums: Dict[ChainSelector, int],
            max_committed_seq_nums: Dict[ChainSelector, int]
    ) -> Dict[ChainSelector, Interval]:
        pass

    # The OCR3 implementation of Outcome
    def outcome(
            self,
            previous_outcome: CommitOutcome,
            query: CommitQuery,
            observations: List[CommitObservation]
    ) -> CommitOutcome:
        match previous_outcome:
            # If the previous round choose an interval, this round we observed merkle roots from the source chains
            # and acquired RMN signed intervals. Use these values to generate a report.
            case NextIntervalsChosen(intervals):
                observed_merkle_roots = self.flatten_merkle_root_observations(observations)
                signed_intervals = self.get_signed_intervals_to_report(intervals, query.signed_intervals,
                                                                       observed_merkle_roots)
                return ReportGenerated(signed_intervals)

            # If the previous round generated a report, this round we checked if the last report has been committed.
            # If it has, output a ReportCommitted value with the new maximum committed sequence numbers (which will
            # be used to construct intervals in the next round). If the report hasn't been committed yet, output a
            # ReportInFlight value so that next round we again check if the report has been committed.
            case ReportGenerated(signed_intervals):
                max_committed_seq_nums = self.flatten_max_committed_seq_nums_observations(observations)
                if self.report_has_been_committed(signed_intervals, max_committed_seq_nums):
                    return ReportCommitted(max_committed_seq_nums)
                else:
                    return ReportInFlight(signed_intervals, 1)

            # If the previous round checked if the report has been committed the dest chain, output a ReportCommitted
            # value if in this round we confirmed that the report has been committed. Else, output a ReportInFlight
            # so that we continue to check next round. If we exhaust attempts, we output
            # ReportCommitted(max_committed_seq_nums) which essentially acts as retrying to build/transmit the report
            # again.
            case ReportInFlight(signed_intervals, attempts):
                max_committed_seq_nums = self.flatten_max_committed_seq_nums_observations(observations)
                if self.report_has_been_committed(signed_intervals, max_committed_seq_nums):
                    return ReportCommitted(max_committed_seq_nums)
                elif attempts >= self.max_check_report_persisted_attempts:
                    return ReportCommitted(max_committed_seq_nums)
                else:
                    return ReportInFlight(signed_intervals, attempts + 1)

            # If in the previous round we detected the report was committed on the dest chain, it means in this round
            # we constructed the intervals for next round. Output NextIntervalsChosen.
            case ReportCommitted(max_committed_seq_nums):
                max_source_chain_seq_nums = self.flatten_max_source_chain_seq_nums_observations(observations)
                next_intervals = self.get_next_intervals(max_source_chain_seq_nums, query.rmn_max_seq_nums,
                                                         max_committed_seq_nums)
                return NextIntervalsChosen(next_intervals)

    # The OCR3 implementation of Report
    def report(self, outcome: CommitOutcome) -> Optional[CommitReport]:
        match outcome:
            case ReportGenerated(signed_intervals):
                return CommitReport(signed_intervals)
            case _:
                return None
