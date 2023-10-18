
(() => {
    // Add click event to fact to show/hide answer
    const answerWrapper = document.querySelectorAll('.answer-wrapper');
    const toggleBtns = document.querySelectorAll('.answer-toggle');

    // Initially, all answers will be hidden
    for (const answerItem of answerWrapper) {
        answerItem.style.display = 'none';
    }

    // Subsequently, the buttons show/hide the answers:
    // if the button is pressed and the answer is hidden, it will show it,
    // or vice versa otherwise
    for (const btn of toggleBtns) {
        btn.addEventListener('click', (e) => {
            const answer = e.target.parentElement.nextElementSibling;
            answer.style.display = answer.style.display === 'none'
                ? 'block'
                : 'none';
        })
    }

    // Handling the update request by the PATCH method using the fetch API
    const editForm = document.querySelector('#form-update-fact')
    // The Id of the fact is obtained from the form dataset by first
    // verifying that it is defined (it would not be if we were on another page)
    const factToEdit = editForm && editForm.dataset.factid;

    editForm && editForm.addEventListener('submit', (e) => {
        e.preventDefault();

        const formData = Object.fromEntries(new FormData(editForm));

        return fetch(`/fact/${factToEdit}`, {
            method: 'PATCH',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData),
        })
            .then(() => document.location.href = `/fact/${factToEdit}`);
    })

    // Handling delete request via DELETE method using the fetch API
    const deleteButton = document.querySelector("#delete-button");
    const factToDelete = deleteButton && deleteButton.dataset.factid;

    deleteButton && deleteButton.addEventListener('click', (e) => {
        const result = confirm("Are you sure you want to delete this fact?")

        if (!result) return;

        return fetch(`/fact/${factToDelete}`, {
            method: 'DELETE',
        })
            .then(() => document.location.href = '/');
    })

})()