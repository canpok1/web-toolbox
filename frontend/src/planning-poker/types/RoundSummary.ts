export interface RoundSummary {
  average: number;
  median: number;
  max: number;
  min: number;
  voteCounts: VoteCount[];
}

export interface VoteCount {
  value: string;
  count: number;
  participants: {
    name: string;
  }[];
}
