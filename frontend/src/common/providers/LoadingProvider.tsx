import type React from "react";
import { type ReactNode, createContext, useState } from "react";

export interface LoadingContextProps {
  showLoading: boolean;
  setShowLoading: (showLoading: boolean) => void;
}

export const LoadingContext = createContext<LoadingContextProps | undefined>(
  undefined,
);

interface LoadingProviderProps {
  children: ReactNode;
}

export const LoadingProvider: React.FC<LoadingProviderProps> = ({
  children,
}) => {
  const [showLoading, setShowLoading] = useState(false);

  const value: LoadingContextProps = {
    showLoading,
    setShowLoading,
  };

  return (
    <LoadingContext.Provider value={value}>{children}</LoadingContext.Provider>
  );
};
