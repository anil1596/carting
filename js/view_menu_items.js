// $(document).ready(function() {
//     //Disable the Remove Button
//     var rowCount = $('table >tbody:last >tr').length;
//     if(rowCount == 1) {
//         document.getElementsByClassName('btn-remove')[0].disabled = true;
//     }
    
//     $(document).on('click', '.btn-add', function(e) {
//         e.preventDefault();
        
//         var controlForm = $('table');
//         var currentEntry = $('table>tbody>tr:last');
//         var newEntry = $(currentEntry.clone()).appendTo(controlForm);
//         // newEntry.find('.VendorMenuItem:firstChild').val('2');  
//         // newEntry.getElementsByClassName("VendorMenuItem")[0].innerHTML="2";
//                                                //Remove the Data - as it is cloned from the above
//         var rowCount = $('table >tbody:last >tr').length;
//         document.getElementsByName("VendorMenuItem")[rowCount-1].lastChild.data=rowCount;
//         //Add the button  
//         var rowCount = $('table >tbody:last >tr').length;
//         if(rowCount > 1) {
//             var removeButtons = document.getElementsByClassName('btn-remove');
//             for(var i = 0; i < removeButtons.length; i++) {
//                 removeButtons.item(i).disabled = false;
//             }
//         }
         
//     }).on('click', '.btn-remove', function(e) {
//         $(this).parents('tr:first').remove();
        
//         //Disable the Remove Button
//         var rowCount = $('table >tbody:last >tr').length;
//         if(rowCount == 1) {
//             document.getElementsByClassName('btn-remove')[0].disabled = true;
//         }

//         e.preventDefault();
//         return false;
//     });
// });

// window.addEventListener('load', function() {
      //To add hotel-names to  drop drown
      function  viewitems(){
      console.log("loading items")


      var vendor_id ={
        "vendorid": 1
      }
 
      fetch('/getvendorsmenu', {
          method: 'POST',
          headers: {
              'Accept': 'application/json',
              'Content-Type': 'application/json'
          },
          credentials: 'same-origin',
          body: JSON.stringify(vendor_id)
      }).then(function(response) {

          return response.json();

      }).then(function(items) {
         
          var displaytable = document.getElementById("view-table-body");
          console.log(displaytable);
          for (var i = 0; i < items.length; i++) {


              var nrow = document.createElement('tr');

              var c1 = document.createElement('td');
              c1.setAttribute('class', "count");
              nrow.appendChild(c1);

              var c2 = document.createElement('td');
              c2.setAttribute('spellcheck', false);
              c2.setAttribute('contenteditable', false);
              c2.innerHTML = items[i]["item_name"];
              nrow.appendChild(c2);

              var c3 = document.createElement('td');
              c3.setAttribute('spellcheck', false);
              c3.setAttribute('contenteditable', false);
              c3.innerHTML = items[i]["price"];
              nrow.appendChild(c3);

              var c4 = document.createElement('td');
              c4.setAttribute('spellcheck', false);
              c4.setAttribute('contenteditable', false);
              c4.innerHTML = items[i]["item_type"];
              nrow.appendChild(c4);

              var c5 = document.createElement('td');
              c5.setAttribute('spellcheck', false);
              c5.setAttribute('contenteditable', false);
              c5.innerHTML = items[i]["item_nature"];
              nrow.appendChild(c5);

              var c6 = document.createElement('td');
              c6.setAttribute('spellcheck', false);
              c6.setAttribute('contenteditable', false);
              c6.innerHTML = items[i]["discount"];
              nrow.appendChild(c6);

              var c7 = document.createElement('td');
              c7.setAttribute('spellcheck', false);
              c7.setAttribute('contenteditable', false);
              c7.innerHTML = items[i]["item_description"];
              nrow.appendChild(c7);

              displaytable.appendChild(nrow);


          }
      }, function(err) {
          console.log(err);
      })
    }
  // })

  