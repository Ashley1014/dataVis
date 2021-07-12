// when user submits form ("On Form Submit Event" snippet):
$("#ticketForm").on("submit", function() {
    let formValid = true;

  // if name component is valid
    if( $("#userName").prop("validity").valid ) {
      // hide name feedback
        $("#userNameFeedback").addClass("hidden");
  // else
    } else {
      // show name feedback
        $("#userNameFeedback").removeClass("hidden");
        formValid = false;
    }

    // if password component is valid
    if( $("#userPassword").prop("validity").valid ) {
        // hide password feedback
        $("#pswFeedback").addClass("hidden");
    // else
    } else {
        // show password feedback
        $("#pswFeedback").removeClass("hidden");
        formValid = false;
    }

    // if confirmed password component is valid
    if( $("#cuserPassword").prop("validity").valid ) {
        // hide confirmed password feedback
        $("#cpswFeedback").addClass("hidden");
    // else
    } else {
        // show confirmed password feedback
        $("#cpswFeedback").removeClass("hidden");
        formValid = false;
    }

    // if confirmed password corresponds to password entered previously
    if( $("#userPassword").val()===$("#cuserPassword").val() ) {
        console.log("inside last if")
        // hide confirmed password feedback
        $("#cpswFeedback2").addClass("hidden");
    } else {
        // show confirmed password feedback
        $("#cpswFeedback2").removeClass("hidden");
        formValid = false;
    }

    // if gender component is valid
    if( $("userGender").val() !== "") {
        // hide gender feedback
        $("#genderFeedback").addClass("hidden");
        // else
    } else {
        // show gender feedback
        $("#genderFeedback").removeClass("hidden");
        formValid = false;
    }

    // if birthday component is valid
    if( $("#userDob").prop("validity").valid ) {
        // hide birthday feedback
        $("#dobFeedback").addClass("hidden");
        // else
    } else {
        // show birthday feedback
        $("#dobFeedback").removeClass("hidden");
        formValid = false;
    }

    // send form to server if formValid is true
    return formValid;
});

$("#loginForm").on("submit", function() {
    let formValid = true;

    // if name component is valid
    if( $("#userName").prop("validity").valid ) {
        // hide name feedback
        $("#userNameFeedback").addClass("hidden");
        // else
    } else {
        // show name feedback
        $("#userNameFeedback").removeClass("hidden");
        formValid = false;
    }

    // if password component is valid
    if( $("#userPassword").prop("validity").valid ) {
        // hide password feedback
        $("#pswFeedback").addClass("hidden");
        // else
    } else {
        // show password feedback
        $("#pswFeedback").removeClass("hidden");
        formValid = false;
    }

    // send form to server if formValid is true
    return formValid;
});
