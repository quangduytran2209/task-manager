import { useState, useEffect } from "react";
import { getTasks, createTask, deleteTask } from "./api";

export default function Dashboard() {
  const [tasks, setTasks] = useState([]);
  const [form, setForm] = useState({ title: "", description: "" });

  useEffect(() => {
    loadTasks();
  }, []);

  async function loadTasks() {
    const data = await getTasks();
    setTasks(data);
  }

  async function handleAdd(e) {
    e.preventDefault();
    await createTask({ ...form, deadline: form.deadline });
    setForm({ title: "", description: "", deadline: "" });
    loadTasks();
  }

  async function handleDelete(id) {
    await deleteTask(id);
    loadTasks();
  }

  return (
    <div className="p-6 max-w-xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">Task Dashboard</h1>

      {/* Add Task Form */}
      <form onSubmit={handleAdd} className="mb-6 space-y-2">
        <input
          className="border p-2 w-full"
          placeholder="Task title"
          value={form.title}
          onChange={(e) => setForm({ ...form, title: e.target.value })}
        />
        <input
          className="border p-2 w-full"
          placeholder="Description"
          value={form.description}
          onChange={(e) => setForm({ ...form, description: e.target.value })}
        />
        <input
          type="date"
          className="border p-2 w-full"
          value={form.deadline}
          onChange={(e) => setForm({ ...form, deadline: e.target.value })}
        />

        <button className="bg-blue-500 text-white px-4 py-2 rounded">
          Add Task
        </button>
      </form>

      {/* Task List */}
      <ul className="space-y-2">
        {tasks.map((t) => (
          <li key={t.id} className="border p-3 rounded flex justify-between">
            <div>
              <h2 className="font-semibold">{t.title}</h2>
              <p className="text-sm text-gray-600">{t.description}</p>
              <p className="text-xs text-gray-400">Status: {t.status}</p>
            </div>
            <button
              onClick={() => handleDelete(t.ID)}
              className="bg-red-500 text-white px-3 py-1 rounded"
            >
              Delete
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}
