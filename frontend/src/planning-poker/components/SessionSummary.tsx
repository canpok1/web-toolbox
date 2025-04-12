import type { Session } from "../types/Session";

function SessionSummary({ session }: { session: Session }) {
  const names = session.participants.map((p) => p.name);
  return (
    <div className="card mx-auto mb-5 max-w-2xl shadow-sm">
      <div className="card-body bg-neutral-content text-left">
        <h2 className="card-title">セッション名: {session.name}</h2>
        <p>参加者: {names.join(", ")}</p>
      </div>
    </div>
  );
}

export default SessionSummary;
