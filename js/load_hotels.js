  $(document).ready(function() {
      $(window).scroll(function() { // check if scroll event happened
          if ($(document).scrollTop() > 50) { // check if user scrolled more than 50 from top of the browser window
              $(".ournav").css("background-color", "rgba(0, 0, 0, 0.7)"); // if yes, then change the color of class "navbar-fixed-top" to white (#f8f8f8)
          } else {
              $(".ournav").css("background-color", "transparent"); // if not, change it back to transparent
          }
      });

  });

  // <!-- <li><a href="#">A9</a></li> -->
  window.addEventListener('load', function() {
      //To add hotel-names to  drop drown
      // console.log("hiii!!");
      var drop = document.getElementsByClassName("hotel-names-list");
      // console.log(drop);
      fetch('/getvendors', {
          method: 'GET',
          headers: {
              'Accept': 'application/json',
              'Content-Type': 'application/json'
          },
          credentials: 'same-origin',
      }).then(function(response) {
          return response.json();
      }).then(function(hotels) {
          for (var i = 0; i < hotels.length; i++) {
              // console.log(hotels[i]["vendor_id"], hotels[i]["vendorname"]);
              var listItem = document.createElement("li");
              listItem.setAttribute("onclick","loadMenuItems("+ hotels[i]["vendor_id"] +")")
              listItem.setAttribute("class","black");             
              listItem.innerHTML = hotels[i]["vendorname"];
              console.log(hotels[i]["vendorname"]);
              drop[0].appendChild(listItem);
          }
      }, function(err) {
          console.log(err);
      })
  })

  //Function to register vendor Called from dashboard




  