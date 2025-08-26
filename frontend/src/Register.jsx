import { useState } from "react";
import { register } from "./api";

export default function Register() {
    const [form, setForm] = useState({ username: "", password: "", email: "" });

    const handleSubmit = async (e) => {
        e.preventDefault();
        const res = await register(form.username, form.password, form.email);
        alert(res.message || res.error);
    };

    return (
        <form onSubmit={handleSubmit} className="p-4 space-y-2 max-w-sm mx-auto">
            <input className="border p-2 w-full" placeholder="Username"
                onChange={(e) => setForm({ ...form, username: e.target.value })} />

            <input className="border p-2 w-full" placeholder="Email"
                onChange={(e) => setForm({ ...form, email: e.target.value })} />

            <input className="border p-2 w-full" type="password" placeholder="Password"
                onChange={(e) => setForm({ ...form, password: e.target.value })} />

            <button className="bg-blue-500 text-white px-4 py-2 rounded">Register</button>
        </form>
    );
}