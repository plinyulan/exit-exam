import { request } from "./index";

export function saveAuth(token: string, role: string) {
  localStorage.setItem("token", token);
  localStorage.setItem("role", role);
}

export function isAdmin(): boolean {
  return localStorage.getItem("role") === "admin";
}


export function login(username: string, password: string) {
  return request("/auth/login", {
    method: "POST",
    body: JSON.stringify({ username, password }),
  });
}
