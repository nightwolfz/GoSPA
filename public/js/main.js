$(document).ready(function(){
	
	$('#contact-button').on('click', function(){
		$('#intro-box').fadeOut(1000, function(){
			$('#contact-box').fadeIn(1000, function(){});
		});
	});
	
});



/*
// Performs a smooth page scroll to an anchor on the same page.
$(function() {
  $('a[href*=#scroll]:not([href=#])').click(function() {
    if (location.pathname.replace(/^\//,'') == this.pathname.replace(/^\//,'') && location.hostname == this.hostname) {
      var target = $(this.hash);
      target = target.length ? target : $('[name=' + this.hash.slice(1) +']');
      if (target.length) {
        $('html,body').animate({
          scrollTop: target.offset().top
        }, 1000);
        return false;
      }
    }
  });
});*/