import VotePanel from "./VotePanel";

export type FibonacciVotePanelProps = {
  roundId: string;
  participantId: string;
  votedOption: string | null;
  onAfterVote: (option: string) => void;
  className?: string;
};

const voteOptions = [
  "0",
  "1",
  "2",
  "3",
  "5",
  "8",
  "13",
  "21",
  "34",
  "55",
  "89",
  "?",
];

export default function FibonacciVotePanel({
  roundId,
  participantId,
  votedOption,
  onAfterVote,
  className,
}: FibonacciVotePanelProps) {
  return (
    <VotePanel
      roundId={roundId}
      participantId={participantId}
      voteOptions={voteOptions}
      votedOption={votedOption}
      onAfterVote={onAfterVote}
      className={className}
    />
  );
}
