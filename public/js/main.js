$(document).ready(function(){
	
	$('#contact-button').on('click', function(){
		$('#intro-box').fadeOut(500, function(){
			//$('#contact-box').fadeIn(1000, function(){
				$('.intro').fadeOut(500, function(){
					$('#scroll-contact').fadeIn(1000);
				});
			//});
		});
	});
	
	$('.brand-heading span:first, .brand-heading span:last, .intro-text, .intro-text + h3').hide();
	
	setTimeout(function(){
		$('.brand-heading span:first').fadeIn(600);
		setTimeout(function(){
			$('.brand-heading span:last').fadeIn(600, function(){
				$('.intro-text').hide().fadeIn(600);
				$('.intro-text + h3').hide().fadeIn(600);
			});
		}, 200);
	}, 200);
	
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