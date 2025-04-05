import type { Participant } from "../types/Participant";

export type SessionSummaryProps = {
  sessionId: string;
  participants: Participant[];
};

function SessionSummary({ sessionId, participants }: SessionSummaryProps) {
  const names = participants.map((p) => p.name);
  return (
    <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
      <div className="card-body bg-neutral-content text-left">
        <h2 className="card-title">セッションID: {sessionId}</h2>
        <p>参加者: {names.join(", ")}</p>
      </div>
    </div>
  );
}

export default SessionSummary;
