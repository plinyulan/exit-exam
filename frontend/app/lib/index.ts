export const API_BASE =
  process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v2";

function getHeaders() {
  const token = typeof window !== "undefined" ? localStorage.getItem("token") : null;
  const role = typeof window !== "undefined" ? localStorage.getItem("role") : null;
  console.log(token);

  return {
    "Content-Type": "application/json",
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
    ...(role ? { "X-ROLE": role } : {}),
  };
}
export async function request(path: string, options: RequestInit = {}) {
  const token = localStorage.getItem("token");
  const role = localStorage.getItem("role");

  const res = await fetch(`${API_BASE}${path}`, {
    ...options,
    // ❌ ลบทิ้ง
    // credentials: "include",
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...(options.headers || {}),
    },
  });

  const data = await res.json();
  if (!res.ok) {
    throw new Error(data.error || "Request failed");
  }
  return data;
}
