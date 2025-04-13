import { useCallback, useState } from "react";
import { ApiClient } from "../../api/ApiClient";
import type { Round } from "../types/Round";

export type ReturnValue = {
  round: Round | null;
  fetch: (roundId: string) => Promise<void>;
};

export default function useRound(): ReturnValue {
  const [round, setRound] = useState<Round | null>(null);

  const fetch = useCallback(async (roundId: string) => {
    const apiClient = new ApiClient();
    const response = await apiClient.fetchRound(roundId);
    if (response.round) {
      setRound({
        ...response.round,
      });
    }
  }, []);
  return {
    round,
    fetch,
  };
}
