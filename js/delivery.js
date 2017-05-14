var rowCount = $('table >tbody:last >tr').length;
        document.getElementsByName("VendorMenuItem")[rowCount-1].lastChild.data=rowCount;