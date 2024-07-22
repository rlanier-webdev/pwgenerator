// Function to request a new password from the server
function generatePassword() {
    fetch('/')
        .then(response => response.text())
        .then(html => {
            // Parse the HTML and extract the new password
            const parser = new DOMParser();
            const doc = parser.parseFromString(html, 'text/html');
            const newPassword = doc.querySelector('#password').value; // Extract the new password value
            document.getElementById('password').value = newPassword; // Update the input field value
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

// Ensure the document is fully loaded before adding event listeners
document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('generate-btn').addEventListener('click', generatePassword);
});
