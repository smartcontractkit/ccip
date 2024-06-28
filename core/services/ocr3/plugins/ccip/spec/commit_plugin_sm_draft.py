"""
High-level Python specification for the CCIPv2 OCR3 Commit Plugin, with a focus on RMN OffChain Blessing

The plugin is implemented as a state machine, and moves from state to state each round. There are 3 states:
1. SelectingIntervalsForReport
      - Determine intervals to be included in the next report
2. BuildingReport
      - Build a report from the intervals determined in the previous round
3. WaitingForReportTransmission
      - Check if the maximum committed sequence numbers on the dest chain have changed since generating the most
        recent report, i.e. check if the report has been committed.
      - If the maximum committed sequence numbers have changed (i.e. the report has been committed) or the maximum
        number of check attempts have been exhausted, move to the SelectingIntervalsForReport state and generate a new
        report.
      - If the maximum committed sequence numbers have _not_ changed (i.e. the report is still in-flight) and the
        maximum number of check attempts are not been exhausted, move to the WaitingForReportTransmission state in order
        to check again.

This approach leads to a clear separation of concerns and addresses the complications that can arise if a report
is not successfully transmitted (as we explicitly only continue once we know the previous report has been committed).
In this design, full messages are no longer in the observations, only merkle roots and intervals are. This reduces the
size of observations, which reduces bandwidth and improves performance.

This is the state machine diagram. States are in boxes, outcomes are within arrows.

              Start
                |
                V
    -------------------------------
    | SelectingIntervalsForReport | <---------|
    -------------------------------           |
                |                             |
        ReportIntervalsSelected               |
                |                             |
                V                             |
        ------------------                    |
        | BuildingReport | -- ReportEmpty --->|
        ------------------                    |
                |                     ReportTransmitted
         ReportGenerated                     or
                |                    ReportNotTransmitted
                V                             |
    --------------------------------          |
    | WaitingForReportTransmission | -------->|
    --------------------------------
            |           ^
            |           |
        ReportNotYetTransmitted
"""

from typing import List, Dict, Optional
from dataclasses import dataclass

RmnNodeId = str
ChainSelector = int

MAX_INTERVAL_LENGTH = 256


@dataclass
class Interval:
    min: int
    max: int


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


# Query data types
##################

@dataclass
class CommitQuery:
    rmn_max_seq_nums: Dict[ChainSelector, int]
    signed_intervals: Dict[ChainSelector, SignedInterval]


# Observation data types
########################

@dataclass
class MerkleRootsObservation:
    merkle_roots: Dict[ChainSelector, bytes]


@dataclass
class SequenceNumbersObservation:
    max_committed_seq_nums: Dict[ChainSelector, int]
    max_uncommitted_seq_nums: Dict[ChainSelector, int]


CommitObservation = SequenceNumbersObservation | MerkleRootsObservation


# Outcome data types
####################

@dataclass
class ReportIntervalsSelected:
    intervals: Dict[ChainSelector, Interval]


@dataclass
class ReportGenerated:
    signed_intervals: Dict[ChainSelector, SignedInterval]


@dataclass
class ReportNotYetTransmitted:
    max_committed_seq_nums: Dict[ChainSelector, int]
    attempts: int


@dataclass
class ReportTransmitted:
    pass


@dataclass
class ReportNotTransmitted:
    pass


@dataclass
class ReportEmpty:
    pass


CommitOutcome = (
        ReportIntervalsSelected |
        ReportGenerated |
        ReportEmpty |
        ReportNotYetTransmitted |
        ReportTransmitted |
        ReportNotTransmitted
)


# State data types
##################

@dataclass
class SelectingIntervalsForReport:
    pass


@dataclass
class BuildingReport:
    intervals: Dict[ChainSelector, Interval]


@dataclass
class WaitingForReportTransmission:
    previous_max_committed_seq_nums: Dict[ChainSelector, int]
    attempts: int


CommitState = SelectingIntervalsForReport | BuildingReport | WaitingForReportTransmission


# Given the outcome of the previous OCR round, return the CommitState of the current round. This effectively
# defines the state transitions of the Commit Plugin.
def current_state(previous_outcome: Optional[CommitOutcome]) -> CommitState:
    match previous_outcome:
        case None:
            return SelectingIntervalsForReport()

        case ReportIntervalsSelected(intervals):
            return BuildingReport(intervals)

        case ReportGenerated(_, max_committed_seq_nums):
            return WaitingForReportTransmission(max_committed_seq_nums, attempts=0)

        case ReportEmpty():
            return SelectingIntervalsForReport()

        case ReportNotYetTransmitted(max_committed_seq_nums, attempts):
            return WaitingForReportTransmission(max_committed_seq_nums, attempts)

        case ReportTransmitted():
            return SelectingIntervalsForReport()

        case ReportNotTransmitted():
            return SelectingIntervalsForReport()


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


@dataclass
class CommitPlugin:
    rmn_client: RmnClient
    all_source_chains: List[ChainSelector]
    dest_chain: ChainSelector
    chain_readers: Dict[ChainSelector, ChainReader]
    f: int
    max_check_report_persisted_attempts: int

    # The OCR3 implementation of Outcome
    def query(self, previous_outcome: Optional[CommitOutcome]) -> CommitQuery:
        match current_state(previous_outcome):
            # If we are choosing the next intervals this round, we need to query RMN for the max uncommitted sequence
            # numbers it has for each source chain, so we can set appropriate upper ranges for our intervals.
            case SelectingIntervalsForReport():
                rmn_max_seq_nums = self.rmn_client.request_max_seq_nums(self.all_source_chains)
                return CommitQuery(rmn_max_seq_nums, {})

            # If we are building a report this round, we request signed intervals from RMN which need to be included
            # in the report
            case BuildingReport(intervals):
                signed_intervals = self.rmn_client.request_signed_intervals(intervals)
                return CommitQuery({}, signed_intervals)

            # If we are checking for an update to the maximum committed sequence numbers this round, we don't need to
            # make a request to RMN
            case WaitingForReportTransmission(_, _):
                return CommitQuery({}, {})

    # Given a mapping from source chains to intervals, for each chain read the messages in its interval and compute
    # the merkle root. Return a mapping from chain to merkle root.
    def get_merkle_roots(self, intervals: Dict[ChainSelector, Interval]) -> Dict[ChainSelector, bytes]:
        pass

    # Read from the dest chain (if possible) and return a mapping from source chain to the maximum sequence number that
    # has been committed to the OffRamp on the dest chain.
    def get_max_committed_seq_nums(self) -> Dict[ChainSelector, int]:
        pass

    # For each source chain, return the maximum sequence number (associated with the dest chain) that is on the
    # source chain's OnRamp
    def get_max_uncommitted_seq_nums(self) -> Dict[ChainSelector, int]:
        pass

    # The OCR3 implementation of Observation
    def observation(self, previous_outcome: Optional[CommitOutcome]) -> CommitObservation:
        match current_state(previous_outcome):
            # If we are choosing the next intervals this round, observe the maximum committed and uncommitted
            # sequence numbers
            case SelectingIntervalsForReport():
                return SequenceNumbersObservation(
                    self.get_max_committed_seq_nums(),
                    self.get_max_uncommitted_seq_nums()
                )

            # If we are building a report this round, we need to observe merkle roots
            case BuildingReport(intervals):
                return MerkleRootsObservation(self.get_merkle_roots(intervals))

            # If we are checking for an update to the maximum committed sequence numbers this round, observe these
            # sequence numbers
            case WaitingForReportTransmission(_, _):
                return SequenceNumbersObservation(self.get_max_committed_seq_nums(), {})

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

    # Given a list of SequenceNumbersObservation, return a flattened consensus on the max committed sequence number
    # for each source chain
    def flatten_max_committed_seq_nums_observations(
            self,
            observations: List[CommitObservation]
    ) -> Dict[ChainSelector, int]:
        pass

    # Given a list of SequenceNumbersObservation and the RMN max uncommitted sequence numbers for each source chain,
    # return the intervals for the next round.
    def choose_next_intervals(
            self,
            observations: List[CommitObservation],
            rmn_max_seq_nums: Dict[ChainSelector, int]
    ) -> Dict[ChainSelector, Interval]:
        pass

    # Return True if the max committed sequence numbers observed this round are different from those observed in a
    # previous round, return False otherwise
    def max_committed_seq_nums_are_updated(self, max_committed_seq_nums, previous_max_committed_seq_nums) -> bool:
        pass

    # The OCR3 implementation of Outcome
    def outcome(
            self,
            previous_outcome: Optional[CommitOutcome],
            query: CommitQuery,
            observations: List[CommitObservation]
    ) -> CommitOutcome:
        match current_state(previous_outcome):
            # If we are choosing the next intervals this round, compute the next intervals using the
            # SequenceNumbersObservations and max uncommitted sequence numbers returned by RMN
            case SelectingIntervalsForReport():
                next_intervals = self.choose_next_intervals(observations, query.rmn_max_seq_nums)
                return ReportIntervalsSelected(next_intervals)

            # If we are building a report this round, we observed merkle roots from the source chains and acquired RMN
            # signed intervals. Use these values to generate a report.
            case BuildingReport(intervals):
                observed_merkle_roots = self.flatten_merkle_root_observations(observations)
                signed_intervals = self.get_signed_intervals_to_report(intervals, query.signed_intervals,
                                                                       observed_merkle_roots)
                if len(signed_intervals) == 0:
                    return ReportEmpty()
                else:
                    return ReportGenerated(signed_intervals)

            # If we are checking if the previously generated report has been transmitted, return ReportTransmitted
            # if an update has been detected, return ReportNotTransmitted if our update checks have been
            # exhausted, or return ReportNotYetTransmitted with an incremented "attempts" value otherwise
            case WaitingForReportTransmission(previous_max_committed_seq_nums, attempts):
                max_committed_seq_nums = self.flatten_max_committed_seq_nums_observations(observations)
                if self.max_committed_seq_nums_are_updated(max_committed_seq_nums, previous_max_committed_seq_nums):
                    return ReportTransmitted()
                elif attempts >= self.max_check_report_persisted_attempts:
                    return ReportNotTransmitted()
                else:
                    return ReportNotYetTransmitted(previous_max_committed_seq_nums, attempts + 1)

    # The OCR3 implementation of Report
    def report(self, outcome: CommitOutcome) -> Optional[CommitReport]:
        match outcome:
            case ReportGenerated(signed_intervals):
                return CommitReport(signed_intervals)
            case _:
                return None
