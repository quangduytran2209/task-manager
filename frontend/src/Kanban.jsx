import { useEffect, useState } from "react";
import { DragDropContext, Droppable, Draggable } from "@hello-pangea/dnd";
import { getTasks, updateTask } from "./api";
import TaskModal from "./components/TaskModal";


const columns = ["todo", "in-progress", "done"];

export default function Kanban() {
    const [tasks, setTasks] = useState([]);
    const [filter, setFilter] = useState("all");
    const [selectedTask, setSelectedTask] = useState(null);

    useEffect(() => {
        loadTasks();
    }, []);

    async function loadTasks() {
        const data = await getTasks();
        setTasks(data);
    }

    async function onDragEnd(result) {
        if (!result.destination) return;

        const { source, destination, draggableId } = result;
        if (source.droppableId !== destination.droppableId) {
            const newStatus = destination.droppableId;
            await updateTask(draggableId, { status: newStatus });
            loadTasks();
        }
    }

    function filteredTasks() {
        if (filter === "overdue") {
            return tasks.filter(
                (t) => t.deadline && new Date(t.deadline) < new Date()
            );
        }
        return tasks;
    }

    return (
        <div className="p-6 max-w-5xl mx-auto">
            <h1 className="text-2xl font-bold mb-6">Kanban Board</h1>

            <div className="flex space-x-3 mb-4">
                <button
                    onClick={() => setFilter("all")}
                    className="px-3 py-1 border rounded"
                >
                    All
                </button>
                <button
                    onClick={() => setFilter("overdue")}
                    className="px-3 py-1 border rounded"
                >
                    Overdue
                </button>
            </div>

            <DragDropContext onDragEnd={onDragEnd}>
                <div className="grid grid-cols-3 gap-6">
                    {columns.map((col) => (
                        <Droppable key={col} droppableId={col}>
                            {(provided) => (
                                <div
                                    {...provided.droppableProps}
                                    ref={provided.innerRef}
                                    className="bg-gray-100 p-4 rounded-lg shadow min-h-[300px]"
                                >
                                    <h2 className="text-lg font-semibold mb-3 capitalize">
                                        {col.replace("-", " ")}
                                    </h2>
                                    {filteredTasks()
                                        .filter((t) => t.status === col)
                                        .map((task, index) => (
                                            <Draggable
                                                key={task.ID.toString()}
                                                draggableId={task.ID.toString()}
                                                index={index}
                                            >
                                                {(provided) => (
                                                    <div
                                                        ref={provided.innerRef}
                                                        {...provided.draggableProps}
                                                        {...provided.dragHandleProps}
                                                        className="bg-white p-3 mb-3 rounded shadow cursor-pointer"
                                                        onClick={() => setSelectedTask(task)}
                                                    >
                                                        <h3 className="font-semibold">{task.title}</h3>
                                                        <p className="text-sm text-gray-600">{task.description}</p>
                                                        {task.deadline && (
                                                            <p className="text-xs text-gray-500">
                                                                Deadline: {new Date(task.deadline).toLocaleDateString()}
                                                            </p>
                                                        )}
                                                    </div>
                                                )}
                                            </Draggable>
                                        ))}
                                    {provided.placeholder}
                                </div>
                            )}
                        </Droppable>
                    ))}
                </div>
            </DragDropContext>
            <TaskModal
                task={selectedTask}
                onClose={() => setSelectedTask(null)}
                onSave={async (form) => {
                    await updateTask(selectedTask.ID, form);
                    setSelectedTask(null);
                    loadTasks();
                }}
            />
        </div>
    );
}
