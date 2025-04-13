import { useCallback, useState } from "react";
import { ApiClient } from "../../api/ApiClient";
import { type Session, isSessionStatus } from "../types/Session";

export type ReturnValue = {
  session: Session | null;
  fetch: (sessionId: string) => Promise<void>;
};

function useSession(): ReturnValue {
  const [session, setSession] = useState<Session | null>(null);

  const fetch = useCallback(async (sessionId: string) => {
    const apiClient = new ApiClient();
    const response = await apiClient.fetchSession(sessionId);
    if (response.session) {
      if (!isSessionStatus(response.session.status)) {
        throw new Error(
          `Invalid session status, status=${response.session.status}`,
        );
      }

      const participants = response.session.participants.map((participant) => {
        return {
          id: participant.participantId,
          name: participant.name,
        };
      });
      setSession({
        sessionId: response.session.sessionId,
        sessionName: response.session.sessionName,
        participants,
        currentRoundId: response.session.currentRoundId,
        hostId: response.session.hostId,
        status: response.session.status,
      });
    }
  }, []);

  return {
    session,
    fetch,
  };
}

export default useSession;
