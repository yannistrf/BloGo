export function saveToken(token) {
    localStorage.setItem("token", token);
}
  
export function getToken() {
    return localStorage.getItem("token");
}
  
export function isLoggedIn() {
    return !!getToken();
}
  
export function doLogout() {
    localStorage.removeItem("token");
}
