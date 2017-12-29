function error_tips(dst_obj, tip_msg) {
    dst_obj.addClass('has-error')
    dst_obj.focus()
    dst_obj.data('toggle', 'tooltip')
    dst_obj.data('placement', 'right')
    dst_obj.attr('title', tip_msg)
    dst_obj.tooltip(
        { container: 'body' }
    )
    dst_obj.tooltip('show')
}

function error_tips_destroy(dst_obj) {
    dst_obj.removeClass('has-error')
    dst_obj.tooltip('destroy')
}

bootstrap_alert = function() {}
bootstrap_alert.success = function(dst_obj, message){ 
    dst_obj.html('<div class="alert alert-success fade in"><a class="close" data-dismiss="alert">×</a><span>'+message+'</span></div>');
}

bootstrap_alert.danger = function(dst_obj, message){
    dst_obj.html('<div class="alert alert-danger fade in"><a class="close" data-dismiss="alert">×</a><span>'+message+'</span></div>')
}