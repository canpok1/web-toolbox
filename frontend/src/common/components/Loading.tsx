import { useLoading } from "../hooks/useLoading";

export default function Loading() {
  const { showLoading } = useLoading();

  return (
    <>
      {showLoading && (
        <div className="fixed top-0 left-0 z-50 flex h-screen w-screen items-center justify-center bg-black bg-opacity-50">
          <div className="h-32 w-32 animate-spin rounded-full border-white border-t-2 border-b-2" />
        </div>
      )}
    </>
  );
}
