import { fetchTasks } from "@/lib/api";
import AddTaskForm from "@/components/AddTaskForm";
import TaskList from "@/components/TaskList";

export default async function Home() {
  const tasks = await fetchTasks();

  return (
    <main className="max-w-xl mx-auto py-12 px-4">
      <h1 className="text-2xl font-bold mb-8 text-gray-800">タスク管理</h1>
      <AddTaskForm />
      <TaskList tasks={tasks} />
    </main>
  );
}
