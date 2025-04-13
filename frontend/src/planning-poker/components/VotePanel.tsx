import { useState } from "react";
import { ApiClient } from "../../api/ApiClient";
import { ExtractErrorMessage } from "../utils/error";
import Alert from "./Alert";

export type VotePanelProps = {
  roundId: string;
  participantId: string;
  voteOptions: string[];
  votedOption: string | null;
  onAfterVote: (option: string) => void;
};

function VotePanel({
  roundId,
  participantId,
  voteOptions,
  votedOption,
  onAfterVote,
}: VotePanelProps) {
  const [errorMessages, setErrorMessages] = useState<string[]>([]);

  const handleClick = async (option: string) => {
    setErrorMessages([]);
    try {
      const client = new ApiClient();
      await client.vote(roundId, {
        participantId,
        value: option,
      });
      onAfterVote(option);
    } catch (error) {
      setErrorMessages(["投票に失敗しました。もう一度お試しください。"]);
      console.error("failed to vote:", error);
    }
  };

  return (
    <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
      <div className="card-body bg-neutral-content text-left">
        <h2 className="card-title">投票</h2>
        <Alert messages={errorMessages} />
        <div className="grid grid-cols-3 gap-4">
          {voteOptions.map((option) => (
            <button
              key={option}
              type="button"
              className={`btn btn-lg ${option === votedOption ? "btn-active btn-accent" : "btn-outline"}`}
              onClick={() => handleClick(option)}
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
