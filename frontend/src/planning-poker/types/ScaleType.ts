const validScales = ["fibonacci", "t-shirt", "power-of-two"] as const;
export type ScaleType = (typeof validScales)[number];

export function isScaleType(value: string): value is ScaleType {
  return validScales.some((v) => v === value);
}
