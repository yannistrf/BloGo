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

export async function getPosts(token, page) {
  const res = await fetch(`${API_BASE}/post/all?page=${page}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return res.json();
}

export async function getMyPosts(token, page) {
  const res = await fetch(`${API_BASE}/user/me?page=${page}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return res.json();
}

export async function createPost(token, title, content) {
  const res = await fetch(`${API_BASE}/post/add`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ title, content }),
  });
  return res.ok;
}

export async function getQueryPosts(token, page, query) {
  const res = await fetch(`${API_BASE}/post/search?query=${query}&page=${page}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return res.json();
}

export async function getPost(token, post_id) {
  const res = await fetch(`${API_BASE}/post/${post_id}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return res.json();
}