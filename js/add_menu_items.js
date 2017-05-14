// var count =0 ;
  //Function to add items to Menu 
  function add_items() {

      var name = document.getElementById('name').value;
      var price = document.getElementById('price').value;
      var discount = document.getElementById('discount').value;
      var nature;
      if (document.getElementById('veg').checked) {
          nature = true;
      } else if (document.getElementById('n_veg').checked) {
          nature = false;
      }
      var Itype;
      if (document.getElementById('starter').checked) {
          Itype = 'starter';
      } else if (document.getElementById('main').checked) {
          Itype = 'main';
      } else if (document.getElementById('desert').checked) {
          Itype = 'desert';
      }
      var description = document.getElementById('description').value;

      var item = {
          "vendor_id": 1,
          "item_name": name,
          "item_type": Itype,
          "item_nature": nature,
          "item_description": description,
          "price": price,
          "imageaddress": "",
          "discount": parseFloat(discount)
      }

      fetch('/additems', {
          method: 'POST',
          headers: {
              'Accept': 'application/json',
              'Content-Type': 'application/json'
          },
          credentials: 'same-origin',
          body: JSON.stringify(item)
      });
      // alert("Form submit successfully.");

      var tableid = document.getElementById("table-body");

      var nrow = document.createElement('tr');

      var c1 = document.createElement('td');
      c1.setAttribute('class', "count");
      nrow.appendChild(c1);

      var c2 = document.createElement('td');
      c2.setAttribute('spellcheck', false);
      c2.setAttribute('contenteditable', true);
      c2.innerHTML = name;
      nrow.appendChild(c2);

      var c3 = document.createElement('td');
      c3.setAttribute('spellcheck', false);
      c3.setAttribute('contenteditable', true);
      c3.innerHTML = price;
      nrow.appendChild(c3);

      var c4 = document.createElement('td');
      c4.setAttribute('spellcheck', false);
      c4.setAttribute('contenteditable', true);
      c4.innerHTML = Itype;
      nrow.appendChild(c4);

      var c5 = document.createElement('td');
      c5.setAttribute('spellcheck', false);
      c5.setAttribute('contenteditable', true);
      c5.innerHTML = nature;
      nrow.appendChild(c5);

      var c6 = document.createElement('td');
      c6.setAttribute('spellcheck', false);
      c6.setAttribute('contenteditable', true);
      c6.innerHTML = discount;
      nrow.appendChild(c6);

      var c7 = document.createElement('td');
      c7.setAttribute('spellcheck', false);
      c7.setAttribute('contenteditable', true);
      c7.innerHTML = description;
      nrow.appendChild(c7);

      tableid.appendChild(nrow);

      document.getElementById('item-form').reset();
      // rows += "<tr><td>" + name + "</td><td>" + gender + "</td><td>" + age + "</td><td>" + city + "</td></tr>";
      // $(rows).appendTo("#vendor_menu_items tbody");
  }
