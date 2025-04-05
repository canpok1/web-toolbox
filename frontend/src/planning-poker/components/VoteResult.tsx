import type { RoundParticipant } from "../types/Participant";

export type VoteResultProps = {
  participant: RoundParticipant;
  revealed: boolean;
};

function VoteResult({ participant, revealed }: VoteResultProps) {
  return (
    <div className="stat place-items-center">
      <div className="stat-title">{participant.name}</div>
      {revealed && <div className="stat-value">{participant.vote ?? ""}</div>}
      {!revealed && (
        <div className="stat-value">{participant.vote ? "済" : "未"}</div>
      )}
    </div>
  );
}

export default VoteResult;
