window.addEventListener("load", function(){
	document.getElementById("menu-list").style.visibility = "hidden";
})

ItemsInCartList = [];

function addItem(dishID) {
	console.log("in add ITem :",dishID);
	var itemCounter = document.getElementById(dishID);
	itemCounter.firstChild.data = +itemCounter.firstChild.data + 1;
	updatePrice(dishID,true );
}

function removeItem(dishID) {
	var itemCounter = document.getElementById(dishID);
	if (+itemCounter.firstChild.data > 1) {
		itemCounter.firstChild.data = +itemCounter.firstChild.data - 1;
		updatePrice(dishID,false );
	}else{
		document.getElementById("key"+dishID).remove();
		var index = ItemsInCartList.indexOf(dishID);
		if (index > -1) {
		    ItemsInCartList.splice(index, 1);
		    updatePrice(dishID,false);
		}
	}
}

function updatePrice (dishID,plus) {
	var itemPriceIndex=0;
	for(var i=0;i<ItemList.length;i++){
		if(dishID == ItemList[i]["dishID"]){
			itemPriceIndex=i;
			break;
		}
	}
	totalAmount = document.getElementById("total-amount");
	if(plus){
		totalAmount.innerHTML = +totalAmount.innerHTML + parseInt(ItemList[itemPriceIndex]["dishPrice"]);
	}else{
		totalAmount.innerHTML = +totalAmount.innerHTML - parseInt(ItemList[itemPriceIndex]["dishPrice"]);
	}
}

function loadMenuItems(vendorid) {
	// Vendorid    int     `json:"vendor_id"`
	// Itemno      int     `json:"item_no,omitempty"`
	// Name        string  `json:"item_name"`
	// IType       string  `json:"item_type"`
	// Nature      bool    `json:"item_nature"`
	// Description string  `json:"item_description"`
	// Price       string  `json:"price"`
	// Image       string  `json:"imageaddress,omitempty"`
	// Discount    float64 `json:"discount,omitempty"`

   console.log("loading items")


      var vendor_id ={
        "vendorid": parseInt(vendorid)
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

      }).then(function(item) {
     
     for (var i = 0; i < item.length; i++) {
        	console.log(item[i]["item_no"],item[i]["item_name"],item[i]["item_type"],item[i]["item_nature"]);
        }   

    document.getElementById("menu-list").style.visibility = "";
    data = {
		items: item
	};
   

	rivets.binders.veg = function(el, value) {
		if (value) {
			el.src = "./img/veg-sticker.png"
		} else {
			el.src = "./img/non-veg-sticker.png"
		}
	}
	rivets.binders.click = function(el, value) {
		el.onclick = function() {
			addToCart(value);
		}
	}

	rivets.bind(document.getElementById("menu-list"), {
		ItemList: data
	});
  });
}


function checkCartItems(dishID) {
	
	if ($.inArray(dishID, ItemsInCartList) == -1) {
		ItemsInCartList.push(dishID);
		return false;
	} else {
		addItem(dishID);
		return true;
	}
}
function loadCartItems(ItemToLoad) {
	if(!checkCartItems(ItemToLoad[0])){
		addElementToCart(ItemToLoad);
		updatePrice(ItemToLoad[0],true);
	}
};

function addToCart(dishID) {
	var dishIndex = 0;
	for (var i = 0; i < ItemList.length; i++) {
		if (ItemList[i]["dishID"] == dishID) {
			dishIndex = i;
			break;
		}
	}
	var item = [];
	item.push(dishID, ItemList[dishIndex]["dishName"], ItemList[dishIndex]["dishPrice"], ItemList[dishIndex]["dishType"]);
	loadCartItems(item);
};

function addElementToCart(ItemToLoad) {
	console.log(ItemToLoad.length);
	cartItems = document.getElementById("cartItems");
	listItem = document.createElement("li");
	listItem.setAttribute("class", "media media-margin");
	listItem.setAttribute("id", "key"+ItemToLoad[0]);

	plusMinusDiv = document.createElement("div");
	plusMinusDiv.setAttribute("class", "media-left btn-group-vertical cart-plus-minus-buttons-group");

	plusBtn = document.createElement("button");
	plusBtn.setAttribute("class", "btn cart-plus-minus-buttons");
	plusBtn.setAttribute("onclick", "addItem('" + ItemToLoad[0] + "');");

	plusIcon = document.createElement("span");
	plusIcon.setAttribute("class", "glyphicon glyphicon-plus");
	plusBtn.appendChild(plusIcon);
	plusMinusDiv.appendChild(plusBtn);

	itemCounter = document.createElement("span");
	itemCounter.setAttribute("class", "item-counter");
	itemCounter.setAttribute("id", ItemToLoad[0]);
	itemCounter.innerHTML = 1;
	plusMinusDiv.appendChild(itemCounter);


	minusBtn = document.createElement("button");
	minusBtn.setAttribute("class", "btn cart-plus-minus-buttons");
	minusBtn.setAttribute("onclick", "removeItem('" + ItemToLoad[0] + "');");

	minusIcon = document.createElement("span");
	minusIcon.setAttribute("class", "glyphicon glyphicon-minus");
	minusBtn.appendChild(minusIcon);
	plusMinusDiv.appendChild(minusBtn);

	listItem.appendChild(plusMinusDiv);
	itemName = document.createElement("div");
	itemName.setAttribute("class", "media-body our-media-body");

	itemHeading = document.createElement("h4");
	itemHeading.setAttribute("class", "media-heading");
	itemHeading.innerHTML = ItemToLoad[1];

	stickerImg = document.createElement("img");
	stickerImg.setAttribute("class", "pull-right item-in-cart-sticker");

	if (ItemToLoad[3]) {
		stickerImg.setAttribute("src", "./img/veg-sticker.png");
	} else {
		stickerImg.setAttribute("src", "./img/non-veg-sticker.png");
	}
	itemHeading.appendChild(stickerImg);

	priceTag = document.createElement("p");
	priceTag.setAttribute("class", "pull-right price-tag item-in-cart-price-tag");
	priceTag.innerHTML = ItemToLoad[2];
	itemHeading.appendChild(priceTag);

	itemName.appendChild(itemHeading);
	listItem.appendChild(itemName);

	cartItems.appendChild(listItem);

	separator = document.createElement("li");
	separator.setAttribute("class", "divider cart-item-separator");
	separator.setAttribute("role", "separator");
	cartItems.appendChild(separator);
}