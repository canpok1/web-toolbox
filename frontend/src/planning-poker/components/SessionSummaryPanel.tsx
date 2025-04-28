import { useMemo } from "react";
import type { Session } from "../types/Session";

export type SessionSummaryPanelProps = {
  session: Session;
  currentParticipantId: string | null;
  className?: string;
};

export default function SessionSummaryPanel({
  session,
  currentParticipantId,
  className,
}: SessionSummaryPanelProps) {
  const currentUserName = useMemo(() => {
    if (!currentParticipantId) {
      return null;
    }
    const currentUser = session.participants.find(
      (p) => p.id === currentParticipantId,
    );
    return currentUser ? currentUser.name : null;
  }, [session.participants, currentParticipantId]);

  const participantNames = session.participants.map((p) => p.name);

  return (
    <div className={`card mx-auto shadow-sm ${className}`}>
      <div className="card-body bg-neutral-content text-left">
        {currentUserName && (
          <p className="font-semibold">あなたの名前: {currentUserName}</p>
        )}
        <p>
          参加者 ({participantNames.length}名):{" "}
          {participantNames.map((p) => (
            <span key={p} className="badge badge-sm">
              {p}
            </span>
          ))}
        </p>
      </div>
    </div>
  );
}
