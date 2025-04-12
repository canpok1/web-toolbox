import { useCallback, useState } from "react";
import { ApiClient } from "../../api/ApiClient";
import type { Session } from "../types/Session";

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
      setSession({
        id: response.session.sessionId,
        name: response.session.sessionName,
        participants: response.session.participants.map((participant) => {
          return {
            id: participant.participantId,
            name: participant.name,
          };
        }),
      });
    }
  }, []);

  return {
    session,
    fetch,
  };
}

export default useSession;
