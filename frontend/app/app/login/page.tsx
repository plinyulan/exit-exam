"use client";

import { SetStateAction, useState } from "react";
import { saveAuth } from "@/lib/auth"

import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"

export default function LoginPage() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");

    async function login() {
        const res = await fetch(process.env.NEXT_PUBLIC_API_URL + "/auth/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ username, password }),
        });

        const data = await res.json();
        if (!res.ok) {
            setError(data.error || "Login failed");
            return;
        }

        saveAuth(data.token, data.role);
        location.href = "/promises";
    }
    return (
        <div className="w-full h-screen flex justify-center items-center">
            <Card className="w-full max-w-sm">
                <CardHeader>
                    <CardTitle>Login to your account</CardTitle>
                </CardHeader>
                <CardContent>
                    <form onSubmit={e => {
                        e.preventDefault();
                        login();
                    }}>
                        <div className="flex flex-col gap-6">
                            <div className="grid gap-2">
                                <Label htmlFor="username">Username</Label>
                                <Input
                                    id="username"
                                    type="text"
                                    onChange={(e: { target: { value: SetStateAction<string>; }; }) => setUsername(e.target.value)}
                                    required
                                />
                            </div>
                            <div className="grid gap-2">
                                <div className="flex items-center">
                                    <Label htmlFor="password">Password</Label>
                                </div>
                                <Input id="password" type="password" onChange={(e: { target: { value: SetStateAction<string>; }; }) => setPassword(e.target.value)} required />
                            </div>
                        </div>
                    </form>
                </CardContent>
                <CardFooter className="flex-col gap-2">
                    <Button type="submit" className="w-full" onClick={login}>
                        Login
                    </Button>
                    <p style={{ color: "red" }}>{error}</p>
                </CardFooter>
            </Card>
        </div>
    )
}

