window.logout = () => {
    fetch('/api/logout', {method: 'POST'}).catch(() => {});
    setTimeout(() => {
        document.cookie = "loginToken=; Max-Age=0; path=/";
        window.location.href = "/login";
    }, 50);
}