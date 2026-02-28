"use client";

import { useTransition } from "react";
import { useRouter } from "next/navigation";

export default function DoneButton({ id }: { id: string }) {
  const [isPending, startTransition] = useTransition();
  const router = useRouter();

  async function handleClick() {
    await fetch(`/api/tasks/${id}/done`, { method: "PUT" });
    startTransition(() => {
      router.refresh();
    });
  }

  return (
    <button
      onClick={handleClick}
      disabled={isPending}
      className="text-xs bg-green-500 text-white px-3 py-1 rounded-md hover:bg-green-600 disabled:opacity-50 transition-colors"
    >
      {isPending ? "..." : "完了"}
    </button>
  );
}
