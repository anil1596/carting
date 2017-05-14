function validate_form(form_id, form_elements_class)
{
	//Get the form-id or the div-id, and search for all elements with the associated class name.
	//Elements that need to be validated need to have a data-regex and a data-alert. data-regex is the
	//attribute that holds the regular expression for validation, and data-alert is the attribute that
	//holds the message that will be placed in the placeholder in case the validation does not take place.
	//Eventually, the functon returns a boolean value - true when the form has been validated, and false when
	//the validation fails.
	var form = document.getElementById(form_id);
	var form_elements = form.getElementsByClassName(form_elements_class);


	var password_array = [];
	var password_array_in_text = [];


	var number_of_elements_to_be_validated = form_elements.length;    //Initialize it to the total number of elements.
	var validated_count = 0;    //Set the number of validated elements to 0 initially.
	for (var i = 0; i < form_elements.length; i++)
	{
		var el = form_elements[i];


		var regex = '';
		try
		{
			regex = new RegExp(el.dataset.regex);
		}
		catch(e)
		{
			//el.dataset.regex does not exist. Hence, the element does not need validation.
			//Reduce the number of elements to be validated by 1.
			number_of_elements_to_be_validated--;
			continue;
		}

		//Remove elements that are passwords and append them to another list.
		if (el.type == "password")
		{
			var temp = form_elements[i];
			password_array.push(temp);
			password_array_in_text.push(temp.value);
			//Reduce the count here, since in case the password is according to the regex, the count will be increased in the next if-else block.
			//This reduction is so that it will neutralize the effect of password validation. The count will be increased again
			//when all the passwords are equal.
			validated_count -= 1;

		}


		var value = el.value;
		var data_alert = el.dataset.alert;
		//Begin the regex match here
		if (regex.test(value))
		{
			el.style.border = "1px solid #ccc";
			el.style.boxShadow = "none";
			//Since the element has been validated, increase the validated_count by 1.
			validated_count += 1;
		}
		else
		{
			el.style.borderColor = "#EF5350";
			el.style.boxShadow = "0px 0px 5px 1px #E0E0E0";
			el.value = '';
			el.placeholder = data_alert;
		}
	}
	//Now, validate passwords
	var unique_passwords = password_array_in_text.filter(function(item, i, ar){ return ar.indexOf(item) === i; });
	//console.log(unique_passwords);
	if (unique_passwords.length == 1)
	{
		//All the passwords entered are the same. Hence, the validation is complete
		//Increment the validated_count by the length of the password_array.
		// console.log("Passwords have matched");
		validated_count += password_array.length;
	}
	else
	{
		//Passwords aren't the same. Highlight the password inputs.
		for (var i = 0; i < password_array.length; i++)
		{
			var el = password_array[i];
			el.style.borderColor = "#EF5350";
			el.style.boxShadow = "0px 0px 5px 1px #E0E0E0";
			el.value = '';
			el.placeholder = el.dataset.alert;
		}
	}

	//Validation has been done. Check how many valid elements are present.
	//If the values match, then the form has been successfully validated. Return a true. Otherwise, return a false.
	if (number_of_elements_to_be_validated == validated_count){return true;}
	else{return false;}
}