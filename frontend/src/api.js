const API_URL="http://localhost:8080/"

export async function register(username, password, email){
    const res = await fetch(`${API_URL}/register`,{
        method: "POST",
        headers: {"Content-Type":"application/json"},
        body: JSON.stringify({username, password, email}),
    });
    return res.json();
}

export async function login(email, password) {
    const res = await fetch(`${API_URL}/login`,{
        method: "POST",
        headers: {"Content-Type":"application/json"},
        body: JSON.stringify({email, password})
    });
    return res.json();
}

function authHeader() {
  return { Authorization: `Bearer ${localStorage.getItem("token")}` };
}

export async function getTasks() {
  const res = await fetch(`${API_URL}/tasks`, { headers: authHeader() });
  return res.json();
}

export async function createTask(task) {
  const res = await fetch(`${API_URL}/tasks`, {
    method: "POST",
    headers: { ...authHeader(), "Content-Type": "application/json" },
    body: JSON.stringify(task),
  });
  return res.json();
}

export async function updateTask(id, task) {
  const res = await fetch(`${API_URL}/tasks/${id}`, {
    method: "PUT",
    headers: { ...authHeader(), "Content-Type": "application/json" },
    body: JSON.stringify(task),
  });
  return res.json();
}

export async function deleteTask(id) {
  const res = await fetch(`${API_URL}/tasks/${id}`, {
    method: "DELETE",
    headers: authHeader(),
  });
  return res.json();
}

await createTask({ ...form, deadline: form.deadline });