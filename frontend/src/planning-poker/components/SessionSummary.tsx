import { useMemo } from "react";
import type { Session } from "../types/Session";

export type SessionSummaryProps = {
  session: Session;
  currentParticipantId: string | null;
  className?: string;
};

function SessionSummary({
  session,
  currentParticipantId,
  className,
}: SessionSummaryProps) {
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
          参加者 ({participantNames.length}名): {participantNames.join(", ")}
        </p>
      </div>
    </div>
  );
}

export default SessionSummary;
