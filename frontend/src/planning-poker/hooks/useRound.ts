import { useCallback, useState } from "react";
import { ApiClient } from "../../api/ApiClient";
import type { Round } from "../types/Round";

export type ReturnValue = {
  round: Round | null;
  fetch: (roundId: string, participantId: string | undefined) => Promise<void>;
};

export default function useRound(): ReturnValue {
  const [round, setRound] = useState<Round | null>(null);

  const fetch = useCallback(
    async (roundId: string, participantId: string | undefined) => {
      const apiClient = new ApiClient();
      const response = await apiClient.fetchRound(roundId, participantId);
      const votes = response.round.votes.map((vote) => {
        return {
          participantId: vote.participantId,
          vote: vote.value ?? null,
        };
      });

      if (response.round) {
        setRound({
          ...response.round,
          votes,
        });
      }
    },
    [],
  );
  return {
    round,
    fetch,
  };
}
