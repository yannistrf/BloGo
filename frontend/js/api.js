const API_BASE = "http://127.0.0.1:8081"; // or your API URL

export async function login(username, password) {
  const res = await fetch(`${API_BASE}/auth/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password }),
  });
  return res.json();
}

export async function register(username, password) {
  const res = await fetch(`${API_BASE}/auth/register`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, password }),
  });
  return res.ok;
}

export async function getPosts(token) {
  const res = await fetch(`${API_BASE}/post/all`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return res.json();
}

export async function getMyPosts(token) {
  const res = await fetch(`${API_BASE}/user/me`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return res.json();
}