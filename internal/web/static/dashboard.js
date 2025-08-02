
document.addEventListener("DOMContentLoaded", () => {
    fetch('/api/status')
        .then(res => res.json())
        .then(data => {
            document.getElementById("status").innerText =
                `Database: ${data.database} | Container: ${data.container} | Time: ${data.timestamp}`;
        });
});

function triggerHealing() {
    fetch('/api/heal', { method: 'POST' })
        .then(res => res.json())
        .then(data => {
            alert(data.status);
        });
}
