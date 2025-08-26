import { useState } from "react";
import { login } from "./api";

export default function Login() {
    const [form, setForm] = useState({ email: "", password: "" });

    const handleSubmit = async (e) => {
        e.preventDefault();
        const res = await login(form.email, form.password);
        if (res.token) {
            localStorage.setItem("token", res.token);
            alert("Login success!");
            window.location.href = "/dashboard";
        } else {
            alert(res.error);
        }
    };

    return (
        <form onSubmit={handleSubmit} className="p-4 space-y-2 max-w-sm mx-auto">
            <input className="border p-2 w-full" placeholder="Email"
                onChange={(e) => setForm({ ...form, email: e.target.value })} />
            <input className="border p-2 w-full" type="password" placeholder="Password"
                onChange={(e) => setForm({ ...form, password: e.target.value })} />
            <button className="bg-green-500 text-white px-4 py-2 rounded">Login</button>
        </form>
    );
}