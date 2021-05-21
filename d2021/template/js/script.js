const ACCEPTED_SECTION_ID = "acceptedSection";
const SUGGESTED_SECTION_ID = "suggestedSection";

function acceptCourse(cardId) {
    let card = document.getElementById(cardId);
    const acceptedSection = document.getElementById(ACCEPTED_SECTION_ID);
    const suggestedSection = document.getElementById(SUGGESTED_SECTION_ID);

   

    acceptedSection.appendChild(card);
    card.querySelector('.cancel-link').removeAttribute("hidden");
    card.querySelector('.enroll-link').setAttribute("hidden", "");
    card.querySelector('.card-subtitle').setAttribute("hidden", "");
    card.querySelector('.progress').removeAttribute("hidden");

}