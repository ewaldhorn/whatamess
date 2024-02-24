class AssemblyLine
  def initialize(speed)
    @speed = speed
  end
  def production_rate_per_hour
    if @speed <= 4
      return @speed * 221.0
    elsif @speed < 9
      return @speed * 221.0 * 0.90
    elsif @speed < 10
      return @speed * 221.0 * 0.80
    else
      return @speed * 221.0 * 0.77
    end
  end
  def working_items_per_minute
    (production_rate_per_hour / 60).floor
  end
end

