import { useState, useEffect } from "react";

export default function TaskModal({ task, onSave, onClose }) {
    const [form, setForm] = useState(task || {});

    useEffect(() => {
        setForm(task || {});
    }, [task]);

    if (!task) return null;

    return (
        <div className="fixed inset-0 bg-black/40 flex items-center justify-center">
            <div className="bg-white rounded-xl shadow-lg p-6 w-96">
                <h2 className="text-xl font-bold mb-4">Edit Task</h2>
                <input
                    className="border p-2 w-full mb-2"
                    value={form.title || ""}
                    onChange={(e) => setForm({ ...form, title: e.target.value })}
                />
                <textarea
                    className="border p-2 w-full mb-2"
                    value={form.description || ""}
                    onChange={(e) => setForm({ ...form, description: e.target.value })}
                />
                <input
                    type="date"
                    className="border p-2 w-full mb-2"
                    value={form.deadline ? form.deadline.split("T")[0] : ""}
                    onChange={(e) => setForm({ ...form, deadline: e.target.value })}
                />

                <div className="flex justify-end gap-2">
                    <button onClick={onClose} className="px-3 py-1 border rounded">
                        Cancel
                    </button>
                    <button
                        onClick={() => onSave(form)}
                        className="px-3 py-1 bg-blue-600 text-white rounded"
                    >
                        Save
                    </button>
                </div>
            </div>
        </div>
    );
}