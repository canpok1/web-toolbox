export type VotePanelProps = {
  voteOptions: string[];
  votedOption: string | null;
  onClick: (option: string) => void;
};

function VotePanel({ voteOptions, votedOption, onClick }: VotePanelProps) {
  return (
    <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
      <div className="card-body bg-neutral-content text-left">
        <h2 className="card-title">投票</h2>
        <div className="grid grid-cols-3 gap-4">
          {voteOptions.map((option) => (
            <button
              key={option}
              type="button"
              className={`btn ${option === votedOption ? "btn-active btn-accent" : "btn-outline"}`}
              onClick={() => onClick(option)}
            >
              {option}
            </button>
          ))}
        </div>
      </div>
    </div>
  );
}

export default VotePanel;
