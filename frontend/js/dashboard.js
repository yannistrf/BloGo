import { getPosts, getMyPosts, createPost, getQueryPosts } from "./api.js";
import { doLogout, getToken } from "./auth.js";

let current_page = 1;
let getPostsFn = getPosts;

const token = getToken();
if (!token) {
  window.location.href = "/login.html";
}

getPosts(token, current_page).then(posts => {
    createPostElements(posts);
})

document.getElementById("newPostForm").addEventListener("submit", async (e) => {
    e.preventDefault();
    const postTitle = document.getElementById("postTitle").value;
    const postContent = document.getElementById("postContent").value;

    const ok = await createPost(token, postTitle, postContent);
    if (ok) {
        document.getElementById("postTitle").value = "";
        document.getElementById("postContent").value = "";
        location.reload();
    } else {
        alert("New post creation failed");
    }
})

document.getElementById("logoutBtn").addEventListener("click", () => {
    doLogout()
    window.location.replace("/login.html")
})

document.getElementById("myPostsBtn").addEventListener("click", () => {
    
    document.getElementById("header").textContent = "Your posts"

    const container = document.getElementById('postsContainer');
    while (container.firstChild) {
        container.removeChild(container.firstChild);
    }

    getPostsFn = getMyPosts;
    current_page = 1;
    document.getElementById("currentPage").textContent = `Page ${current_page}`;


    getMyPosts(token, current_page).then(posts => {
        createPostElements(posts);
    });
})

document.getElementById("nextBtn").addEventListener("click", () => {
    
    const container = document.getElementById('postsContainer');
    
    const query = document.getElementById("query").value;   // if it exists

    getPostsFn(token, current_page + 1, query).then(posts => {
        if (posts.length == 0) {
            alert("There aren't any more posts")
            return;
        }
        
        while (container.firstChild) {
            container.removeChild(container.firstChild);
        }
        createPostElements(posts);
        current_page++;
        document.getElementById("currentPage").textContent = `Page ${current_page}`;
    });
})

document.getElementById("prevBtn").addEventListener("click", () => {
    
    if (current_page === 1)
        return;

    const query = document.getElementById("query").value; // if it exists

    getPostsFn(token, current_page - 1, query).then(posts => {
        const container = document.getElementById('postsContainer');
        while (container.firstChild) {
            container.removeChild(container.firstChild);
        }
        createPostElements(posts);
        current_page--;
        document.getElementById("currentPage").textContent = `Page ${current_page}`
    });

})

document.getElementById("queryForm").addEventListener("submit", async (e) => {
    e.preventDefault()
    const query = document.getElementById("query").value;
    if (query === "") {
        return
    }
    document.getElementById("header").textContent = `Search results for: ${query}`

    const container = document.getElementById('postsContainer');
    while (container.firstChild) {
        container.removeChild(container.firstChild);
    }

    getPostsFn = getQueryPosts;
    current_page = 1;
    document.getElementById("currentPage").textContent = `Page ${current_page}`;

    getQueryPosts(token, current_page, query).then(posts => {
        createPostElements(posts);
    });
})

function createPostElements(posts) {
    const container = document.getElementById('postsContainer');
    posts.forEach(post => {
        const date = new Date(post.created_at)
        const dateString = date.toLocaleDateString();
        const timeString = date.toLocaleTimeString();

        const postElement = document.createElement('div');
        postElement.innerHTML = `
        <a href="/post.html?post_id=${post.id}" style="text-decoration: none;">
            <div class="card shadow-sm" style="max-width: 800px;">
                <div class="card-body">
                    <div class="d-flex justify-content-between flex-wrap">
                        <h5 class="card-title">${post.title}</h5>
                        <p class="text-muted">Written by: ${post.author}</p>
                    </div>
                    <p class="card-text" style="white-space: pre-wrap;">${post.content}</p>
                    <p class="text-muted mb-0">Created at: ${dateString} (${timeString})</p>
                </div>
            </div>
        </a>
        `;
        container.appendChild(postElement);
    });
}

