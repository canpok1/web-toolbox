import { CircleCheckBig, CircleSlash } from "lucide-react";
import type { RoundParticipant } from "../types/Participant";

export type VoteResultProps = {
  participant: RoundParticipant;
  revealed: boolean;
};

function VoteResult({ participant, revealed }: VoteResultProps) {
  const makeTooltipMessage = () => {
    if (participant.vote === null) {
      return "未投票";
    }
    return revealed ? participant.vote : "投票済み";
  };

  return (
    <div className="tooltip" data-tip={makeTooltipMessage()}>
      <div className="stats h-full w-full shadow">
        <div className="stat place-items-center">
          <div className="stat-title">{participant.name}</div>
          {!participant.isVoted && (
            <CircleSlash className="text-error" aria-label="未投票" />
          )}
          {participant.isVoted && revealed && (
            <div className="stat-value">{participant.vote}</div>
          )}
          {participant.isVoted && !revealed && (
            <CircleCheckBig
              className="text-success-content"
              aria-label="投票済"
            />
          )}
        </div>
      </div>
    </div>
  );
}

export default VoteResult;
