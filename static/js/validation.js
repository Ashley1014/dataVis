console.log("loading...");
// when user submits form ("On Form Submit Event" snippet):
$("#ticketForm").on("submit", function() {
    let formValid = true;

    // TODO: place all "check component validity snippets" here
  // if name component is valid ("Check Component Validity" snippet):
    if( $("#userName").prop("validity").valid ) {
      // hide name feedback
        $("#userNameFeedback").addClass("hidden");
  // else
    } else {
      // show name feedback
        $("#userNameFeedback").removeClass("hidden");
        formValid = false;
    }

    // TODO: snippet(s) for if component data is valid
    // if email component is valid ("Check Component Validity" snippet):
    if( $("#userPassword").prop("validity").valid ) {
        // hide email feedback
        $("#pswFeedback").addClass("hidden");
    // else
    } else {
        // show email feedback
        $("#pswFeedback").removeClass("hidden");
        formValid = false;
    }

    // if numTickets component is valid ("Check Component Validity" snippet):
    if( $("#cuserPassword").prop("validity").valid ) {
        // hide Tickets feedback
        $("#cpswFeedback").addClass("hidden");
    // else
    } else {
        // show numTickets feedback
        $("#cpswFeedback").removeClass("hidden");
        formValid = false;
    }

    //console.log("i'm here!")

    if( $("#userPassword").val()==$("#cuserPassword").val() ) {
        console.log("inside last if")
        $("#cpswFeedback2").addClass("hidden");
    } else {
        $("#cpswFeedback2").removeClass("hidden");
        formValid = false;
    }

    // send form to server if formValid is true (included as part of "On Form Submit Event" snippet)
    return formValid;
});

$("#loginForm").on("submit", function() {
    let formValid = true;

    // TODO: place all "check component validity snippets" here
    // if name component is valid ("Check Component Validity" snippet):
    if( $("#userName").prop("validity").valid ) {
        // hide name feedback
        $("#userNameFeedback").addClass("hidden");
        // else
    } else {
        // show name feedback
        $("#userNameFeedback").removeClass("hidden");
        formValid = false;
    }

    // TODO: snippet(s) for if component data is valid
    // if email component is valid ("Check Component Validity" snippet):
    if( $("#userPassword").prop("validity").valid ) {
        // hide email feedback
        $("#pswFeedback").addClass("hidden");
        // else
    } else {
        // show email feedback
        $("#pswFeedback").removeClass("hidden");
        formValid = false;
    }

    // send form to server if formValid is true (included as part of "On Form Submit Event" snippet)
    return formValid;
});
