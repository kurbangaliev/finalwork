async function logout() {
    await fetch('/logout', {
        method: 'POST',
        credentials: 'include'
    });
    window.location.href = '/';
}