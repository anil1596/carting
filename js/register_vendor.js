  // document.addEventListener("load", function() {
  //     $(window).scroll(function() { // check if scroll event happened
  //         if ($(document).scrollTop() > 50) { // check if user scrolled more than 50 from top of the browser window
  //             $(".ournav").css("background-color", "rgba(0, 0, 0, 0.7)"); // if yes, then change the color of class "navbar-fixed-top" to white (#f8f8f8)
  //         } else {
  //             $(".ournav").css("background-color", "transparent"); // if not, change it back to transparent
  //         }
  //     });

  // });



   //Function to register vendor Called from dashboard
  function register_vendor() {
    console.log("pressed");


    var validated = validate_form("VendorProfile", "__form-validate");
    if(!validated){
      return;
    }

    var name = document.getElementById("vendor_name");
    var email = document.getElementById("vendor_email");
    var first = document.getElementById("owner_first_name");
    var last_name = document.getElementById("owner_last_name");
    var mobile = document.getElementById("vendor_mobile");
    var addr = document.getElementById("vendor_address1");
    var city = document.getElementById("vendor_address_city");
    var country = document.getElementById("vendor_address_country");
    var pin = document.getElementById("vendor_address_postalcode");
    var description = document.getElementById("vendor_description");
    var offers = document.getElementById("vendor_offers");
    var vendor_name = document.getElementById("vendor_name");

    if(name.validity.valueMissing || first.validity.valueMissing || last_name.validity.valueMissing){
      alert("Enter Name fields.");
    }
    if(mobile.validity.valueMissing){
      alert("Enter mobile number");
    }
    if(!mobile.validity.valueMissing){
      if(mobile.value.length != 10){
        alert("Enter 10-digit mobile number");
      }
    }
    if(!email.validity.valid){
      alert("Enter valid email address.");
    }
    var VendorPic = localStorage.getItem('VendorPic');
    
    if(VendorPic == null){
      VendorPic = "";
    }
    
    if (!name.validity.valueMissing &&
      email.validity.valid &&
      !first.validity.valueMissing &&
      !last_name.validity.valueMissing &&
      !mobile.validity.valueMissing &&
      !description.validity.valueMissing) {
      
      var msg = {
        "owner": first.value + " " + last_name.value,
        "vendorname": name.value,
        "email": email.value,
        "mobile": [mobile.value],
        "address": addr.value + " " + city.value + " " + country.value + " " + pin.value,
        "imageaddress": VendorPic,
        "description": description.value,
        "offer": offers.value,
        "password": "desitadka123"
      }
      fetch('/registervendor', {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        credentials: 'same-origin',
        body: JSON.stringify(msg)
       }); //.then(function(response){
      //       return response.json();
      // }).then(function(vid){
      //    console.log("vid": vid["vendorid"]);

      // });
      console.log(VendorPic);
      // alert("Form submit successfully.");
    }
  }


   function previewFile(){
       var preview = document.querySelector('img'); //selects the query named img
       var file    = document.querySelector('input[type=file]').files[0]; //sames as here
       var reader  = new FileReader();

       reader.onloadend = function () {
           preview.src = reader.result;
           console.log(reader.result);
           window.open(reader.result);
           var vendorPic = {
            vendorPicture : reader.result
           }
           localStorage.setItem('vendorPic', vendorPic);
       }

       if (file) {
           reader.readAsDataURL(file); //reads the data as a URL
       } else {
           preview.src = "https://ununsplash.imgix.net/photo-1431578500526-4d9613015464?fit=crop&fm=jpg&h=300&q=75&w=400";
       }
  }