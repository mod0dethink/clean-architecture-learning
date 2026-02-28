export type TaskStatus = "open" | "done";

export interface Task {
  ID: string;
  Title: string;
  Status: TaskStatus;
  CreatedAt: string;
}

// Server Components からバックエンドに直接アクセスする（サーバー間通信）
const BACKEND_URL =
  process.env.BACKEND_URL ?? "http://localhost:8080";

export async function fetchTasks(): Promise<Task[]> {
  const res = await fetch(`${BACKEND_URL}/api/tasks`, {
    cache: "no-store",
  });
  if (!res.ok) throw new Error("Failed to fetch tasks");
  return res.json();
}
