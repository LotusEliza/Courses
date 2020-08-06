$(window).on("load",function(){
    setTimeout(function () {
        $('#hellopreloader').fadeOut('slow',function(){$(this).remove()});
    }, 400);

    $('.sliderBox__info').addClass('fadeInLeft vis');
});
/*preload*/

$(document).ready(function() {
    Core.init();
});

var Core = function() {


    var runHeader = function() {
        // menu scroll to section
        class StickyNavigation {
            constructor() {
                this.currentId = null;
                this.currentTab = null;
                this.tabContainerHeight = 100;
                let self = this;
                $('.js-scroll-trigger').click(function(event) {
                    event.preventDefault();
                    var title = $(this).attr('data-name');
                    //console.log("block: " + title);
                    //gtag('event', 'screen_view', {'screen_name' : title});
                    gtag('event', 'view_item', {'event_label' : title});
                    self.onTabClick(event, $(this));
                });
                $(window).scroll(() => { this.onScroll(); });
            }

            onTabClick(event, element) {
                event.preventDefault();
                let scrollTop = $(element.attr('href')).offset().top - this.tabContainerHeight + 1;
                $('html, body').animate({ scrollTop: scrollTop }, 600);
            }

            onScroll() {
                this.checkTabContainerPosition();
                this.findCurrentTabSelector();
            }

            checkTabContainerPosition() {
                let offset = $('.sliderBox').offset().top + $('.sliderBox').height() - this.tabContainerHeight;
                if($(window).scrollTop() > offset) {
                    $('.header').addClass('header-sticky');
                }
                else {
                    $('.header').removeClass('header-sticky');
                }
            }

            findCurrentTabSelector(element) {
                let newCurrentId;
                let newCurrentTab;
                let self = this;
                $('.js-scroll-trigger').each(function() {
                    let id = $(this).attr('href');
                    let offsetTop = $(id).offset().top - self.tabContainerHeight;
                    let offsetBottom = $(id).offset().top + $(id).height() - self.tabContainerHeight;
                    if($(window).scrollTop() > offsetTop && $(window).scrollTop() < offsetBottom) {
                        newCurrentId = id;
                        newCurrentTab = $(this);
                    }
                });
                if(this.currentId != newCurrentId || this.currentId === null) {
                    this.currentId = newCurrentId;
                    this.currentTab = newCurrentTab;
                    $(".js-scroll-trigger").removeClass('active')
                    $(".js-scroll-trigger[href='"+newCurrentId+"']").addClass('active')
                    if( newCurrentId == '#aboutUs'){
                    }
                }
            }
        }
        new StickyNavigation();

        // adaptive menu
        $('.menuBtnToggle').on('click',function(e) {
            e.preventDefault();
            $(this).toggleClass('active');
            if($(this).hasClass('active')) {
                $('.header').addClass('mobileHeader');
                $('.navBox').addClass('mobileActive');
                $('.page-overlay').fadeIn(300);

            }
            else {
                $('.header').removeClass('mobileHeader');
                $('.navBox').removeClass('mobileActive');
                $('.page-overlay').fadeOut(300);

            }
        });
        $('.page-overlay').on('click',function() {
            $('.menuBtnToggle').toggleClass('active');
            $('.navBox').toggleClass('mobileActive');
            $('.page-overlay').fadeToggle(300);
        });
        /////
    }

    var runContent = function() {
        $('#modal').iziModal();
        $('.bttn-course').on('click',function(event) {
            event.preventDefault();
            var title = $(this).attr('data-name');
            gtag('event', 'view_item', {'event_label' : title});

            $('#modal').iziModal('open');
            $('.hidden-course-name').val(title);
            $('.modal-record .title-container .text').html(title);
            $('.form-record').removeClass('dn');
            $('.success-form').addClass('dn');
            $('.error-form').addClass('dn');
            /*$('#modal .form').append('<input type="hidden" class="hidden-course-name" name="курс" value="'+ title + '"> ')*/
        });


        $('#form-record').on('submit', function(e) {

            var $form = $(this);
            e.preventDefault();
            $.ajax({
                type: 'POST',
                dataType: 'json',
                url: $form.attr('action'),
                data: $form.serialize(),
                success: function(data) {
                    console.log('success');
                    console.log(data);
                    $form.trigger('reset');
                    $('.form-record').addClass('dn');
                    $('.success-form').removeClass('dn');
                },

                error: function(data) {
                    console.log('error');
                    console.log(data);
                    $('.form-record').addClass('dn');
                    $('.success-form').addClass('dn');
                    $('.error-form').removeClass('dn');
                },
            });
        });

        //animation
        var animationSetting = function(){
            let winHeight = $(window).height();
            let scroll = $(window).scrollTop();
            let offset1 = $('#aboutUs').offset().top;
            let offset2 = $('#orderOfTraining').offset().top;
            let offset3 = $('#courses').offset().top;
            let offset4 = $('#counters').offset().top;
            let offset5 = $('#map').offset().top;
            if (scroll > offset1-(winHeight/1.25)) {
                $('#aboutUs .image-container').addClass('fadeInLeft vis');
                $('#aboutUs .text-container').addClass('fadeInRight vis');
            }
            if (scroll > offset2-(winHeight/1.25)) {
                $('#orderOfTraining .step').addClass('zoomIn  vis');
            }
            if (scroll > offset3-(winHeight/1.5)) {
                $('#courses .course').addClass('flipInY vis');
            }
            if (scroll > offset4-(winHeight)) {
                $('#counters .counter').addClass('rotateIn vis');
            }
            if (scroll > offset5-(winHeight/1.25)) {
                $('#map .map-info').addClass('fadeInRight vis');
            }
        }
        /*animationSetting();*/
        $(document).scroll(animationSetting);
    }
    var initSliders = function(){
        if ($("#mainSlider")[0]) {
            $('#mainSlider').slick({
                autoplay: true,
                autoplaySpeed: 6000,
                slidesToShow: 1,
                fade:true,
                arrows:true,
                dots: false,
                draggable:false,
                responsive: [
                    {
                        breakpoint: 992,
                        settings: {
                            draggable:true,
                        }
                    }
                ]

            });
        }
        if ($("#reviews-slider")[0]) {
            $('#reviews-slider').slick({
                autoplay: false,
                slidesToShow: 2,
                arrows: false,
                dots: true,
                infinite: false,
                responsive: [
                    {
                        breakpoint: 992,
                        settings: {
                            slidesToShow: 2,
                            slidesToScroll: 1
                        }
                    },
                    {
                        breakpoint: 768,
                        settings: {
                            slidesToShow: 1,
                            slidesToScroll: 1
                        }
                    }
                ]
            });
        }
    }



    return {
        init: function() {
          /*  runHeader();
            runContent();*/
            runHeader();
            runContent();
            initSliders();
        }
    }
}()















