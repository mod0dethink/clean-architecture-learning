import type { Task } from "@/lib/api";
import DoneButton from "./DoneButton";

export default function TaskList({ tasks }: { tasks: Task[] }) {
  if (tasks.length === 0) {
    return (
      <p className="text-gray-400 mt-6 text-center">タスクがありません。</p>
    );
  }

  return (
    <ul className="mt-6 space-y-2">
      {tasks.map((task) => (
        <li
          key={task.ID}
          className="flex items-center justify-between p-4 bg-white border border-gray-200 rounded-lg shadow-sm"
        >
          <span
            className={
              task.Status === "done"
                ? "line-through text-gray-400"
                : "text-gray-800"
            }
          >
            {task.Title}
          </span>
          {task.Status === "open" && <DoneButton id={task.ID} />}
          {task.Status === "done" && (
            <span className="text-xs text-green-500 font-medium">完了</span>
          )}
        </li>
      ))}
    </ul>
  );
}
