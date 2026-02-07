import { request } from "./index";

// GET /promises/all
export function getAllPromises() {
  return request("/promises/all");
}

// GET /promises/{id}
export function getPromiseDetail(id: number) {
  return request(`/promises/${id}`);
}

// POST /promises/{id}/updates
export function addPromiseUpdate(
  id: number,
  note: string,
  updatedAt = new Date().toISOString()
) {
  return request(`/promises/${id}/updates`, {
    method: "POST",
    body: JSON.stringify({
      note,
      updated_at: updatedAt,
    }),
  });
}
