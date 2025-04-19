import { CircleCheckBig, CircleSlash } from "lucide-react";
import type { Vote } from "../types/Vote";

export type VoteResultProps = {
  name: string;
  vote?: Vote;
  revealed: boolean;
};

function VoteResult({ name, vote, revealed }: VoteResultProps) {
  // 公開時は未投票の情報は表示しない
  if (revealed && !vote) {
    return null;
  }

  const makeTooltipMessage = () => {
    if (!vote) {
      return "未投票";
    }
    return revealed ? vote.vote : "投票済み";
  };

  return (
    <div className="tooltip" data-tip={makeTooltipMessage()}>
      <div className="stats h-full w-full shadow">
        <div className="stat place-items-center">
          <div className="stat-title">{name}</div>
          {!vote && <CircleSlash className="text-error" aria-label="未投票" />}
          {vote && revealed && <div className="stat-value">{vote.vote}</div>}
          {vote && !revealed && (
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
